package indodana

// Response Get Installment Available
type ResponseCheckoutPayment struct {
	Status        string `json:"status"`
	RedirectURL   string `json:"redirectUrl,omitempty"`
	TransactionID string `json:"transactionId,omitempty"`
	Error         Error  `json:"error,omitempty"`
}

// Response Get Installment Available
type ResponseGetInstallment struct {
	Status  string    `json:"status"`
	Payment []Payment `json:"payments,omitempty"`
	Error   Error     `json:"error,omitempty"`
}

// Response Check Status Transaction
type ResponseCheckStatus struct {
	Status            string  `json:"status"`
	FraudStatus       string  `json:"fraudStatus"`
	LegalName         string  `json:"legalName"`
	Amount            float64 `json:"amount"`
	PaymentType       string  `json:"paymentType"`
	TransactionStatus string  `json:"transactionStatus"`
	MerchantOrderID   string  `json:"merchantOrderId"`
	TransactionTime   string  `json:"transactionTime"`
	TransactionID     string  `json:"transactionId"`
	EntityID          string  `json:"entityId,omitempty"`
	Message           string  `json:"message,omitempty"`
}

// Error general message f
type Error struct {
	Kind    string      `json:"kind"`
	Message interface{} `json:"message"`
}

// Response For Installment Calculation
type Payment struct {
	ID                           string  `json:"id"`
	Tenure                       int     `json:"tenure"`
	Rate                         float64 `json:"rate"`
	Amount                       float64 `json:"amount"`
	PaymentType                  string  `json:"paymentType"`
	MonthlyInstallment           float64 `json:"monthlyInstallment"`
	MerchantDiscountRate         float64 `json:"merchantDiscountRate,omitempty"`
	DiscountedMonthlyInstallment float64 `json:"discountedMonthlyInstallment,omitempty"`
	InstallmentAmount            float64 `json:"installmentAmount"`
	DownPayment                  float64 `json:"downPayment,omitempty"`
	Provision                    float64 `json:"provision,omitempty"`
}
