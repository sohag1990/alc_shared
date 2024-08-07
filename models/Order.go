package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Property struct {
	DefaultProperties
	LandlordName     string
	LandlordPhone    string
	LandlordEmail    string
	PropertyCategory string //commercial or residential
	PropertyType     string // if commercial: shop, office, resturount etc other then HMO, flat, house, communal area
	PropertyAddress  string
	PropertyPostCode string
	FlatNumber       string
	UserID           uint64
	Orders           []Order
	OrderServices    []OrderServices
}

// Service public model generated by bindu
type Order struct {
	DefaultProperties
	OrderServices       []OrderServices
	UserID              uint64
	PropertyID          uint64
	PropertyFullAddress string
	Property            Property
	Subtotal            float32
	Discount            float32
	Total               float32
	LandlordName        string
	FullName            string
	Email               string
	Phone               string
	CompanyName         string
	BillingAddress      string
	BillingPostCode     string
	OrderNotes          string
	InvID               uint64
	DeletedAt           *time.Time            `sql:"index" json:"-"`
	IsDel               soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	OrderStatus         string                // same as woocommerce order status
	PaymentStatus       string
	PaymentMethod       string
	CreatedBy           string
	AppointmentTimeSlot string
	AppointmentDate     string
	Hostname            string
	OrderDate           *time.Time
	Archived            bool
}

// OrderService public model generated by bindu
type OrderServices struct {
	OrderID    uint64
	ServiceID  int
	PropertyID uint64
	Title      string
	Subtotal   float32
	Discount   float32
	Total      float32
	ID         uint64 `gorm:"primary_key"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time            `json:"-"`
	DeletedAt  *time.Time            `sql:"index" json:"-"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	Index      int
	// extra added for testing
	ReviewDate time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	ReminderSentDate time.Time `form:"ReminderSentDate"  time_format:"2006-01-02"`
	RenewalStatus    bool
	RenewedDate      time.Time `form:"RenewedDate"  time_format:"2006-01-02"`
	RenewdOrderID    uint64
	RemainingDays    int
}
