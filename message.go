package infobip

import "regexp"

// BulkMessage contains the body request for multiple messages
type BulkMessage struct {
	ID       string    `json:"bulkId,omitempty"`
	Messages []Message `json:"messages"`
}

// Message contains the body request
type Message struct {
	From            string        `json:"from,omitempty"`
	Destinations    []Destination `json:"destinations,omitempty"`
	To              string        `json:"to,omitempty"`
	Text            string        `json:"text"`
	Transliteration string        `json:"transliteration,omitempty"`
	LanguageCode    string        `json:"languageCode,omitempty"`
	NotifyURL       string        `json:"notifyUrl,omitempty"`
}

// Destination contains the recipient
type Destination struct {
	ID string `json:"messageId,omitempty"`
	To string `json:"to"`
}

// MessageResponse body response
type MessageResponse struct {
	BulkID   string        `json:"bulkId,omitempty"`
	Messages []MessageInfo `json:"messages"`
}

// MessageInfo ...
type MessageInfo struct {
	ID       string        `json:"messageId"`
	To       string        `json:"to"`
	Status   MessageStatus `json:"status"`
	SMSCount int           `json:"smsCount"`
}

// MessageStatus ...
type MessageStatus struct {
	ID          int    `json:"id"`
	Action      string `json:"action,omitempty"`
	GroupID     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MessageStatusWithID ...
type MessageStatusWithID struct {
	BulkID string        `json:"bulkId,omitempty"`
	ID     string        `json:"messageId"`
	To     string        `json:"to"`
	Status MessageStatus `json:"status"`
}

// MessageStatusResponse ...
type MessageStatusResponse struct {
	Results []MessageStatusWithID `json:"results"`
}

// Validate validates the entire message values
func (b BulkMessage) Validate() (err error) {
	for _, m := range b.Messages {
		if err = m.Validate(); err != nil {
			break
		}
	}
	return
}

// Validate validates the body request values
func (m Message) Validate() (err error) {
	if err = m.validateFromValue(); err != nil {
		return
	}
	if err = m.validateDestinationValues(); err != nil {
		return
	}
	err = m.validateToValue()
	return
}

func (m Message) validateFromValue() (err error) {
	if isNumeric(m.From) && !isValidRange(m.From, 3, 14) {
		err = ErrForFromNonAlphanumeric
		return
	}
	if !isValidRange(m.From, 3, 13) {
		err = ErrForFromAlphanumeric
		return
	}
	return
}

func (m Message) validateDestinationValues() (err error) {
	for _, d := range m.Destinations {
		if isNumeric(d.To) && !isValidRange(d.To, 3, 14) {
			err = ErrForDestinationNonAlphanumeric
			break
		}
	}
	return
}

func (m Message) validateToValue() (err error) {
	if m.To == "" {
		return
	}
	if isNumeric(m.To) && !isValidRange(m.To, 3, 14) {
		err = ErrForToNonAlphanumeric
		return
	}
	return
}

func isNumeric(s string) bool {
	return regexp.MustCompile(`^[\d]*$`).MatchString(s)
}

func isValidRange(s string, a, b int) bool {
	l := len(s)
	return l > a && l <= b
}
