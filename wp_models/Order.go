package wp_models

type Billing struct {
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	Company   string `json:"company"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Postcode  string `json:"postcode"`
	State     string `json:"state"`
}

type Shipping struct {
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	Company   string `json:"company"`
	Country   string `json:"country"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Postcode  string `json:"postcode"`
	State     string `json:"state"`
}

type LineItem struct {
	ID          int           `json:"id"`
	Image       Image         `json:"image"`
	MetaData    []interface{} `json:"meta_data"` // Empty array in the provided JSON
	Name        string        `json:"name"`
	ParentName  string        `json:"parent_name"`
	Price       float64       `json:"price"`
	ProductID   int           `json:"product_id"`
	Quantity    int           `json:"quantity"`
	SKU         string        `json:"sku"`
	Subtotal    string        `json:"subtotal"`
	SubtotalTax string        `json:"subtotal_tax"`
	TaxClass    string        `json:"tax_class"`
	Taxes       []interface{} `json:"taxes"` // Empty array in the provided JSON
	Total       string        `json:"total"`
	TotalTax    string        `json:"total_tax"`
	VariationID int           `json:"variation_id"`
}

type Image struct {
	ID  string `json:"id"`
	Src string `json:"src"`
}

type MetaDataItem struct {
	ID    int         `json:"id"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"` // Could be a string, number,
}
type Order struct {
	ID                 int            `json:"id"`
	CustomerID         int            `json:"customer_id"`
	Billing            Billing        `json:"billing"`
	Shipping           Shipping       `json:"shipping"`
	LineItems          []LineItem     `json:"line_items"`
	MetaData           []MetaDataItem `json:"meta_data"`
	NeedsPayment       bool           `json:"needs_payment"`
	NeedsProcessing    bool           `json:"needs_processing"`
	Number             string         `json:"number"`
	OrderKey           string         `json:"order_key"`
	ParentID           int            `json:"parent_id"`
	PaymentMethod      string         `json:"payment_method"`
	PaymentMethodTitle string         `json:"payment_method_title"`
	PaymentURL         string         `json:"payment_url"`
	PricesIncludeTax   bool           `json:"prices_include_tax"`
	Refunds            []interface{}  `json:"refunds"`        // Empty array in the provided JSON
	ShippingLines      []interface{}  `json:"shipping_lines"` // Empty array in the provided JSON
	ShippingTax        string         `json:"shipping_tax"`
	ShippingTotal      string         `json:"shipping_total"`
	Status             string         `json:"status"`
	TaxLines           []interface{}  `json:"tax_lines"` // Empty array in the provided JSON
	Total              string         `json:"total"`
	TotalTax           string         `json:"total_tax"`
	TransactionID      string         `json:"transaction_id"`
	Version            string         `json:"version"`
}
