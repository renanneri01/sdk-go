package order

type Request struct {
	Type              string           `json:"type"`
	TotalAmount       string           `json:"total_amount"`
	ExternalReference string           `json:"external_reference"`
	CaptureMode       string           `json:"capture_mode,omitempty"`
	ProcessingMode    string           `json:"processing_mode,omitempty"`
	Description       string           `json:"description,omitempty"`
	Marketplace       string           `json:"marketplace,omitempty"`
	MarketPlaceFee    string           `json:"marketplace_fee,omitempty"`
	ExpirationTime    string           `json:"expiration_time,omitempty"`
	IntegrationData   *IntegrationData `json:"integration_data,omitempty"`
	Transactions      Transaction      `json:"transactions"`
	Payer             Payer            `json:"payer"`
	Items             []Items          `json:"items,omitempty"`
}

type IntegrationData struct {
	IntegrationID string   `json:"integration_id,omitempty"`
	PlatformID    string   `json:"platform_id,omitempty"`
	Sponsor       *Sponsor `json:"sponsor,omitempty"`
}

type Sponsor struct {
	ID *string `json:"id,omitempty"`
}

type Transaction struct {
	Payments []Payment `json:"payments"`
}

type Payment struct {
	Amount        string        `json:"amount"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}

type PaymentMethod struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Token               string `json:"token"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments"`
}

type Payer struct {
	Email          string          `json:"email"`
	FirstName      string          `json:"first_name,omitempty"`
	LastName       string          `json:"last_name,omitempty"`
	Identification *Identification `json:"identification,omitempty"`
	Phone          *Phone          `json:"phone,omitempty"`
}

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type Items struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UnitPrice   string `json:"unit_price"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
	Type        string `json:"type"`
	PictureUrl  string `json:"picture_url"`
	Quantity    int    `json:"quantity"`
}
