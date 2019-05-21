package infobip

type SearchNumberParmas struct {
	Number       string `json:"capabilities,omitempty" url:"capabilities,omitempty"`
	Capabilities string `json:"capabilities,omitempty" url:"capabilities,omitempty"`
	Country      string `json:"country,omitempty" url:"country,omitempty"`
	Limit        int    `json:"limit,omitempty" url:"limit,omitempty"`
	Page         int    `page:"limit,omitempty" url:"page,omitempty"`
}

type SearchNumberResponse struct {
	NumberCount int      `json:"numberCount,omitempty"`
	Numbers     []Number `json:"numbers,omitempty"`
}

type Number struct {
	NumberKey    string      `json:"numberKey"`
	Number       string      `json:"number"`
	Country      string      `json:"country"`
	Type         string      `json:"type,omitempty"`
	Capabilities []string    `json:"capabilities"`
	Shared       bool        `json:"shared,omitempty"`
	Price        NumberPrice `json:"price,omitempty"`
}

type NumberPrice struct {
	PricePerMonth     float32 `json:"pricePerMonth,omitempty"`
	SetupPrice        float32 `json:"setupPrice,omitempty"`
	InitialMonthPrice float32 `json:"initialMonthPrice,omitempty"`
	Currency          string  `json:"currency,omitempty"`
}
