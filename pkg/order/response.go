package order

type Response struct {
	ID                string              `json:"id"`
	Type              string              `json:"type"`
	TotalAmount       string              `json:"total_amount"`
	ExternalReference string              `json:"external_reference"`
	CountryCode       string              `json:"country_code"`
	Status            string              `json:"status"`
	StatusDetail      string              `json:"status_detail"`
	CaptureMode       string              `json:"capture_mode"`
	ProcessingMode    string              `json:"processing_mode"`
	Description       string              `json:"description,omitempty"`
	Marketplace       string              `json:"marketplace,omitempty"`
	MarketplaceFee    string              `json:"marketplace_fee,omitempty"`
	ExpirationTime    string              `json:"expiration_time,omitempty"`
	CreatedDate       string              `json:"created_date"`
	LastUpdatedDate   string              `json:"last_updated_date"`
	Transactions      TransactionResponse `json:"transactions"`
	Payer             Payer               `json:"payer"`
	Items             []Items             `json:"items,omitempty"`
}

type TransactionResponse struct {
	Payments []PaymentResponse `json:"payments"`
	Refunds  []Refund          `json:"refunds,omitempty"`
}

type PaymentResponse struct {
	ID            string        `json:"id"`
	ReferenceID   string        `json:"reference_id"`
	Status        string        `json:"status"`
	Amount        string        `json:"amount"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}

type Refund struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
}

type RefundReference struct {
	ID       string `json:"id"`
	SourceID string `json:"source_id"`
}
