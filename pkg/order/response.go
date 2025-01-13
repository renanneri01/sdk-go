package order

type Response struct {
	ID                string                  `json:"id"`
	ProcessingMode    string                  `json:"processing_mode"`
	ExternalReference string                  `json:"external_reference"`
	Description       string                  `json:"description"`
	Marketplace       string                  `json:"marketplace"`
	MarketplaceFee    string                  `json:"marketplace_fee"`
	TotalAmount       string                  `json:"total_amount"`
	ExpirationTime    string                  `json:"expiration_time"`
	SiteID            string                  `json:"site_id"`
	UserID            string                  `json:"user_id"`
	CreatedDate       string                  `json:"created_date"`
	LastUpdatedDate   string                  `json:"last_updated_date"`
	Type              string                  `json:"type"`
	Status            string                  `json:"status"`
	StatusDetail      string                  `json:"status_detail"`
	CaptureMode       string                  `json:"capture_mode"`
	IntegrationData   IntegrationDataResponse `json:"integration_data"`
	Transactions      TransactionResponse     `json:"transactions"`
	Payer             Payer                   `json:"payer"`
	Items             []Items                 `json:"items"`
}

type IntegrationDataResponse struct {
	ApplicationID string          `json:"application_id"`
	IntegratorID  string          `json:"integrator_id"`
	PlatformID    string          `json:"platform_id"`
	Sponsor       SponsorResponse `json:"sponsor"`
}

type SponsorResponse struct {
	ID string `json:"id"`
}

type TransactionResponse struct {
	Payments []PaymentResponse `json:"payments"`
	Refunds  []Refund          `json:"refunds"`
}

type PaymentResponse struct {
	ID            string        `json:"id"`
	Amount        string        `json:"amount"`
	Status        string        `json:"status"`
	StatusDetail  string        `json:"status_detail"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	Reference     Reference     `json:"reference"`
}

type Reference struct {
	ID       string   `json:"id"`
	Source   string   `json:"source"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
}

type Refund struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Amount        string    `json:"amount"`
	Reference     Reference `json:"reference"`
}

type RefundReference struct {
	ID       string `json:"id"`
	SourceID string `json:"source_id"`
}
