package order

type Request struct {
	Type              string      `json:"type"`
	TotalAmount       string      `json:"total_amount"`
	ExternalReference string      `json:"external_reference"`
	CaptureMode       string      `json:"capture_mode,omitempty"`
	ProcessingMode    string      `json:"processing_mode,omitempty"`
	Description       string      `json:"description,omitempty"`
	Marketplace       string      `json:"marketplace,omitempty"`
	MarketPlaceFee    string      `json:"marketplace_fee,omitempty"`
	ExpirationTime    string      `json:"expiration_time,omitempty"`
	Transactions      Transaction `json:"transactions"`
	Payer             Payer       `json:"payer"`
	Items             []Items     `json:"items,omitempty"`
}

type Transaction struct {
	Payments []Payment `json:"payments"`
}

type Payment struct {
	Amount            string            `json:"amount"`
	PaymentMethod     PaymentMethod     `json:"payment_method"`
	AutomaticPayments *AutomaticPayment `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredential `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionData `json:"subscription_data,omitempty"`
}

type PaymentMethod struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Token               string `json:"token"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments"`
}

type AutomaticPayment struct {
	PaymentProfileID string `json:"payment_profile_id"`
	ScheduleDate     string `json:"schedule_date"`
	DueDate          string `json:"due_date"`
	Retries          int    `json:"retries"`
}

type StoredCredential struct {
	PaymentInitiator   string `json:"payment_initiator"`
	Reason             string `json:"reason"`
	StorePaymentMethod bool   `json:"store_payment_method"`
	FirstPayment       bool   `json:"first_payment"`
}

type SubscriptionData struct {
	InvoiceID            string               `json:"invoice_id"`
	BillingDate          string               `json:"billing_date"`
	SubscriptionSequence SubscriptionSequence `json:"subscription_sequence"`
	InvoicePeriod        InvoicePeriod        `json:"invoice_period"`
}

type SubscriptionSequence struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

type InvoicePeriod struct {
	Type   string `json:"type"`
	Period int    `json:"period"`
}

type Payer struct {
	Email          string          `json:"email"`
	FirstName      string          `json:"first_name,omitempty"`
	LastName       string          `json:"last_name,omitempty"`
	CustomerID     *string         `json:"customer_id,omitempty"`
	Identification *Identification `json:"identification,omitempty"`
	Phone          *Phone          `json:"phone,omitempty"`
	Address        *Address        `json:"address,omitempty"`
}

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type Address struct {
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
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
