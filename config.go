package infobip

const (
	ForwardTypePull = "PULL"
	ForwardTypePost = "HTTP_FORWARD_POST"
	ForwardTypeGet  = "HTTP_FORWARD_GET"
)

type ConfigResponse struct {
	ConfigurationKey string `json:"configurationKey,omitempty" url:"configurationKey,omitempty"`
	IsActive         bool   `json:"isActive,omitempty" url:"isActive,omitempty"`
}

type Action struct {
	ActionKey    string `json:"actionKey,omitempty" url:"actionKey,omitempty"`
	Type         string `json:"type,omitempty" url:"type,omitempty"`
	ForwardURL   string `json:"forwardUrl,omitempty" url:"forwardUrl,omitempty"`
	CallbackData string `json:"callbackData,omitempty" url:"callbackData,omitempty"`
}
