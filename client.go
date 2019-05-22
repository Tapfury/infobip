package infobip

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	//SingleMessagePath for sending a single message
	SingleMessagePath = "sms/1/text/single"

	//AdvancedMessagePath for sending advanced messages
	AdvancedMessagePath = "sms/1/text/advanced"

	// AvailableNumberPath for searching number
	AvailableNumberPath = "numbers/1/numbers/available"

	// RentNumberPath for purchasing number
	RentNumberPath = "numbers/1/numbers"

	// SMSStatusPath for getting status
	SMSStatusPath = "sms/1/reports"
)

// HTTPInterface helps Infobip tests
type HTTPInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client manages requests to Infobip
type Client struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient HTTPInterface
}

// ClientWithBasicAuth returns a pointer to infobip.Client with Infobip funcs
func ClientWithBasicAuth(username, password string) *Client {
	return &Client{
		BaseURL:    "https://api.infobip.com/",
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}
}

// Rent return a newly purchased number
func (c Client) Rent(numberKey string) (*Number, error) {
	data := map[string]string{
		"numberKey": numberKey,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	num := &Number{}

	if err := c.defaultPostRequest(b, RentNumberPath, num); err != nil {
		return nil, err
	}

	return num, nil
}

// SearchNumber return a list of available number
func (c Client) SearchNumber(parmas SearchNumberParmas) (*SearchNumberResponse, error) {
	v, err := query.Values(parmas)
	if err != nil {
		return nil, err
	}

	path := AvailableNumberPath + "?" + v.Encode()
	resp := &SearchNumberResponse{}

	if err := c.defaultGetRequest(path, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// SingleMessage sends one message to one recipient
func (c Client) SingleMessage(m Message) (*MessageResponse, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	resp := &MessageResponse{}
	if err := c.defaultPostRequest(b, SingleMessagePath, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// AdvancedMessage sends messages to the recipients
func (c Client) AdvancedMessage(m BulkMessage) (*MessageResponse, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	resp := &MessageResponse{}
	if err := c.defaultPostRequest(b, AdvancedMessagePath, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSMSStatus return sms status
func (c Client) GetSMSStatus(messageID string) (*MessageStatusWithID, error) {
	params := url.Values{}
	params.Add("messageId", messageID)

	path := SMSStatusPath + "?" + params.Encode()

	resp := &MessageStatusResponse{}
	if err := c.defaultGetRequest(path, resp); err != nil {
		return nil, err
	}

	if len(resp.Results) != 1 {
		return nil, ErrSMSStatusNotFound
	}

	return &resp.Results[0], nil
}

func (c Client) defaultPostRequest(b []byte, path string, v interface{}) error {
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+path, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c Client) defaultGetRequest(path string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}
