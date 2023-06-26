package models

import (
	"time"
)

type Company struct {
	DefaultProperties
	Name                   string
	Address                string
	Phone                  string
	Mobile                 string
	Email                  string
	SiteURL                string
	SSRURL                 string
	BackendTitle           string
	FrontendTitle          string
	ShortTitle             string
	SmallLogo              string
	LargeLogo              string
	Copyright              string
	Keyword                string
	MetaDescription        string
	CurrencySign           string
	CurrencyCode           string
	DefaultLanguage        string
	LicenseKey             string
	SecretKey              string
	SiteOffline            string
	OfflineMessage         string
	AllowRegistration      string
	BookingCancellation    string
	DefaultShiftName       string
	OneDayDeductionForLate int

	Weekend        string
	Branches       []Branch
	Projects       []Project
	Department     []Department
	Options        []Option
	Sections       []Section
	Divisions      []Division
	Areas          []Area
	Districts      []District
	TerritoryAreas []TerritoryArea
	FiscalYears    []FiscalYear
	Shifts         []Shift
	Modules        []Module
	Employees      []Employee
	Holidays       []Holiday
	ODs            []OD
}
type Branch struct {
	DefaultProperties
	CompanyID uint64
	Company   Company

	Name      string
	Address   string
	Phone     string
	Mobile    string
	Email     string
	Employees []Employee
	Projects  []Project
	Holidays  []Holiday
	ODs       []OD
}
type Project struct {
	DefaultProperties
	CompanyID   uint64
	Company     Company
	BranchID    uint64
	Branch      Branch
	Name        string
	Description string
	Price       float64
	Budget      float64
}
type Department struct {
	DefaultProperties
	CompanyID        uint64
	Company          Company
	Name             string
	HeadOfDepartment string
	Status           string
	Sections         []Section
	ODs              []OD
}
type Section struct {
	DefaultProperties
	CompanyID       uint64
	Company         Company
	DepartmentID    uint64
	Department      Department
	Name            string
	SectionHeadName string
	Status          string
	ODs             []OD
}
type Division struct {
	DefaultProperties
	CompanyID      uint64
	Company        Company
	Name           string
	Status         string
	Areas          []Area
	TerritoryAreas []TerritoryArea
	Modules        []Module
}
type Area struct {
	DefaultProperties
	CompanyID      uint64
	Company        Company
	DivisionID     uint64
	Division       Division
	Name           string
	Status         string
	TerritoryAreas []TerritoryArea
}
type District struct {
	DefaultProperties
	CompanyID uint64
	Company   Company
	Name      string
	Status    string
}
type TerritoryArea struct {
	DefaultProperties
	CompanyID  uint64
	Company    Company
	DivisionID uint64
	Division   Division
	AreaID     uint64
	Area       Area
	Name       string
	Status     string
}
type FiscalYear struct {
	DefaultProperties
	CompanyID uint64
	Company   Company
	Year      string
	Status    string
	YearStart string
	YearEnd   string
	Holidays  []Holiday
	ODs       []OD
}

type Shift struct {
	DefaultProperties
	CompanyID     uint64
	Company       Company
	Name          string
	Status        string
	StartTime     string
	EndTime       string
	ShiftCapacity int
	ODs           []OD
}
type Module struct {
	DefaultProperties
	CompanyID   uint64
	Company     Company
	Name        string
	Division    Division
	DivisionID  uint64
	OrderNumber int
	Status      string
	Options     []Option
}

type Option struct {
	DefaultProperties
	CompanyID  uint64
	Company    Company
	Module     Module
	ModuleID   uint64
	MenuName   string
	OptionType string
	OptionName string
	Status     string
}

type Employee struct {
	DefaultProperties
	Attendances []Attendance
	ODs         []OD
	Customers   []Customer
	Weekend     string
	// appointment information
	Company         Company
	CompanyID       uint64
	Branch          Branch
	BranchID        uint64
	Department      Department
	DepartmentID    uint64
	Section         Section
	SectionID       uint64
	Shift           Shift
	ShiftID         uint64
	AppointmentType string
	AppointmentDate time.Time `gorm:"type:date;default:null"`
	JoiningDate     time.Time `gorm:"type:date;default:null"`
	EPRID           uint64
	EmployeeID      uint64
	EmployeeType    string
	Designation     string
	EmployeeImage   string

	// employee information
	EmployeeNameBangla     string
	EmployeeNameEnglish    string
	FatherName             string
	MotherName             string
	HusbandOrWifeName      string
	DateOfBirth            time.Time `gorm:"type:date;default:null"`
	PresentAddress         string
	PermanentAddress       string
	EducationQualification string
	ExtraQualification     string
	Nationality            string
	Gender                 string
	BloodGroup             string
	MaritalStatus          string
	Religion               string
	Mobile                 string
	Phone                  string
	Email                  string

	// salary structer
	CashSalary           float64
	TAndTAllowance       float64
	OthersPayable        float64
	TotalFixedPayable    float64
	BasicSalary          float64
	HouseRentAllowance   float64
	MedicalAllowance     float64
	ConveyanceAllowance  float64
	FestivalBonusPercent float64
	OthersAllowance      float64
	GrossSalary          float64
	ProvidentFund        float64
	IncomeTaxPercent     float64
	IncomeTaxTK          float64
	TotalLoanAndAdvance  float64
	MonthlyInstallment   float64
	TotalLoanPaid        float64
	GrossDeduction       float64
	NetSalary            float64
	PHHeadMapping        string
	LoanHeadMapping      string
	Status               string
}

type Holiday struct {
	DefaultProperties
	CompanyID    uint64
	Company      Company
	BranchID     uint64
	Branch       Branch
	FiscalYearID uint64
	FiscalYear   FiscalYear
	Name         string
	FromDate     string
	ToDate       string
	HolidayImage string
	Status       string
}
type Attendance struct {
	DefaultProperties

	Employee   Employee
	EmployeeID uint64
	Mobile     string
	InTime     string
	OutTime    string
	Remarks    string
	IsPresent  bool
}

type OD struct {
	DefaultProperties
	Employee     Employee
	EmployeeID   uint64
	Company      Company
	CompanyID    uint64
	Branch       Branch
	BranchID     uint64
	Department   Department
	DepartmentID uint64
	Section      Section
	SectionID    uint64
	FiscalYear   FiscalYear
	FiscalYearID uint64
	Shift        Shift
	ShiftID      uint64
	Remarks      string
	ODTo         time.Time `gorm:"type:date;default:null"`
	ODFrom       time.Time `gorm:"type:date;default:null"`
	InTime       time.Time `gorm:"type:date;default:null"`
	OutTime      time.Time `gorm:"type:date;default:null"`
}

type Customer struct {
	DefaultProperties
	//Customer Details
	Name            string
	ShortCode       string
	Address         string
	ShippingAddress string
	//Contact Person Information
	ContactPerson string
	Designation   string
	Gender        string
	Nationality   string
	Employee      Employee
	EmployeeID    uint64
	Mobile        string
	Phone         string
	Fax           string
	Email         string
}

type Distributor struct {
	DefaultProperties
	Name            string
	ShortCode       string
	Address         string
	BillingAddress  string
	ShippingAddress string
	//Contact Person Information
	ContactPerson string
	Designation   string
	Gender        string
	Currency      string
	Attention     string
	Mobile        string
	Phone         string
	Fax           string
	Email         string
}
type Importer struct {
	DefaultProperties
	Name            string
	ShortCode       string
	Address         string
	BillingAddress  string
	ShippingAddress string
	ContactPerson   string
	Designation     string
	Currency        string
	Mobile          string
	Phone           string
	Fax             string
	Email           string
}

// type Workorder struct {
// 	DefaultProperties
// 	Customer          string
// 	WorkorderNo       string
// 	WorkorderType     string
// 	WorkorderDate     string
// 	SalesPerson       string
// 	OEM               string
// 	DeliveryDate      string
// 	IncludingVATAIT   bool
// 	VATPercentage     float64
// 	AITPercentage     float64
// 	WorkorderProducts []WorkorderProduct

// 	TotalWordorder        float64
//     DiscountPercentage   float64
//     DiscountAmount       float64
//     SubTotal             float64
//     PaymentMode          string
//     VATPercentage        float64
//     VATAmount            float64
//     GrandTotal           float64
//     MidmanCommission     float64
//     AITPercentage        float64
//     AITAmount            float64
//     NetReceivable        float64
//     WorkorderAttachment  string
//     CustomPaymentTerms   string
//     Notes                string
// }

//	type WorkorderProduct struct {
//		DefaultProperties
//		ProductCategory    ProductCategory
//		ProductCategoryID  uint64
//		ProductDescription string
//		ProductSKU         string
//		Validity           string
//		Quantity           int
//		UnitPrice          float64
//		Total              float64
//		Remarks            string
//		Workorder          Workorder
//		WorkorderID        uint64
//	}
type ProductCategory struct {
	DefaultProperties
	Name string
}
