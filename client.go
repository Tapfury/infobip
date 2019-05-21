package infobip

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	//SingleMessagePath for sending a single message
	SingleMessagePath = "sms/1/text/single"

	//AdvancedMessagePath for sending advanced messages
	AdvancedMessagePath = "sms/1/text/advanced"
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
