package indodana

// This object use for get available installment
type RequestNotification struct {
	Amount            float64 `json:"amount"`
	PaymentType       string  `json:"paymentType"`
	TransactionStatus string  `json:"transactionStatus"`
	MerchantOrderId   string  `json:"merchantOrderId"`
	TransactionTime   string  `json:"transactionTime"`
	TransactionID     string  `json:"transactionId"`
}

// This object use for get available installment
type RequestGetInstallment struct {
	Amount float64 `json:"amount"`
	Items  []Item  `json:"items"`
}

// This object use checkout payment
type RequestCheckoutPayment struct {
	TransactionDetail       TransactionDetail `json:"transactionDetails"`
	CustomerDetail          CustomerDetail    `json:"customerDetails"`
	Seller                  []Seller          `json:"sellers"`
	BillingAddress          *Address          `json:"billingAddress"`
	ShippingAddress         *Address          `json:"shippingAddress"`
	PaymentType             string            `json:"paymentType"`
	ApprovedNotificationURL string            `json:"approvedNotificationUrl"`
	CancellationRedirectURL string            `json:"cancellationRedirectUrl"`
	BackToStoreURL          string            `json:"backToStoreUrl"`
	ExpirationAt            string            `json:"expirationAt"`
}

// This object use for get status at indodana
type RequestCheckStatus struct {
	MerchantOrderId string `url:"merchantOrderId"`
}

// This object use at payload checkout payment
type TransactionDetail struct {
	Amount          float64 `json:"amount"`
	MerchantOrderId string  `json:"merchantOrderId"`
	Items           []Item  `json:"items"`
}

// This object use at payload get available installment, and checkout payment
type Item struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Url        string  `json:"url,omitempty"`
	ImageURL   string  `json:"imageUrl,omitempty"`
	Type       string  `json:"type,omitempty"`
	Category   string  `json:"category,omitempty"`
	Quantity   int64   `json:"quantity"`
	ParentType string  `json:"parentType"`
	ParentID   string  `json:"parentId"`
}

// This object use at payload checkout payment
type CustomerDetail struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// This object use at payload checkout payment
type Seller struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	URL            string   `json:"url"`
	SellerIDNumber string   `json:"sellerIdNumber"`
	Email          string   `json:"email"`
	Address        *Address `json:"address"`
}

// This object use at seller object
type Address struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
}

type Logger struct {
	Method       string      `json:"method"`
	Host         string      `json:"host"`
	Path         string      `json:"path"`
	RequestBody  interface{} `json:"request_body"`
	ResponseBody interface{} `json:"response_body"`
	TimeLength   int         `json:"time_length"`
}
