package models

// Profile public model generated by bindu
type Profile struct {
	FullName    string
	CompanyName string
	Address     string
	Postcode    string
	Phone       string
	UserID      uint64 `gorm:"unique;not null"`
	DefaultProperties
}

type DashboardReport struct {
	TotalCertificates int64

	ExpiringProperty          []Property
	TotalProperties           int64
	Properties                []Property
	Complaince                int64
	TotalOrders               int64
	RecentOrders              []Order
	ExpiringCertificates      []OrderServices
	OrdersHasExpiringServices []Order
	TotalExpiringCertificates int64
	TotalCustomers            int64
}
