package models

import "time"

type Company struct {
	DefaultProperties
	Name                   string `gorm:"size:250"`
	Address                string `gorm:"size:512"`
	Phone                  string `gorm:"size:20"`
	Mobile                 string `gorm:"size:100"`
	Email                  string `gorm:"uniqueIndex;size:256;not null"`
	SiteURL                string `gorm:"size:350"`
	SSRURL                 string `gorm:"size:512"`
	BackendTitle           string `gorm:"size:512"`
	FrontendTitle          string `gorm:"size:512"`
	ShortTitle             string `gorm:"size:512"`
	SmallLogo              string `gorm:"size:512"`
	LargeLogo              string `gorm:"size:512"`
	Copyright              string `gorm:"size:250"`
	Keyword                string `gorm:"size:1000"`
	MetaDescription        string `gorm:"size:1000"`
	CurrencySign           string `gorm:"size:25"`
	CurrencyCode           string `gorm:"size:25"`
	DefaultLanguage        string `gorm:"size:10"`
	LicenseKey             string `gorm:"size:512"`
	SecretKey              string `gorm:"size:512"`
	SiteOffline            string `gorm:"size:25"`
	OfflineMessage         string `gorm:"size:512"`
	AllowRegistration      string `gorm:"size:25"`
	BookingCancellation    string `gorm:"size:25"`
	DefaultShiftName       string `gorm:"size:50"`
	OneDayDeductionForLate int
	Weekend                string `gorm:"size:250"`

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
	Menu           []Menu
	Employees      []Employee
	Holidays       []Holiday
	ODs            []OD
	Leaves         []Leave
}
type Branch struct {
	DefaultProperties
	CompanyID uint64
	Company   Company

	Name      string `gorm:"size:250"`
	Address   string `gorm:"size:512"`
	Phone     string `gorm:"size:25"`
	Mobile    string `gorm:"size:25"`
	Email     string `gorm:"size:250"`
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
	Name        string `gorm:"size:250"`
	Description string `gorm:"size:512"`
	Price       float64
	Budget      float64
}
type Department struct {
	DefaultProperties
	CompanyID        uint64
	Company          Company
	Name             string `gorm:"size:250"`
	HeadOfDepartment string `gorm:"size:250"`
	Status           string `gorm:"size:25"`
	Sections         []Section
	ODs              []OD
}
type Section struct {
	DefaultProperties
	CompanyID       uint64
	Company         Company
	DepartmentID    uint64
	Department      Department
	Name            string `gorm:"size:250"`
	SectionHeadName string `gorm:"size:250"`
	Status          string `gorm:"size:25"`
	ODs             []OD
}
type Division struct {
	DefaultProperties
	CompanyID      uint64
	Company        Company
	Name           string `gorm:"size:250"`
	Status         string `gorm:"size:25"`
	Areas          []Area
	TerritoryAreas []TerritoryArea
}
type Area struct {
	DefaultProperties
	CompanyID      uint64
	Company        Company
	DivisionID     uint64
	Division       Division
	Name           string `gorm:"size:250"`
	Status         string `gorm:"size:25"`
	TerritoryAreas []TerritoryArea
}
type District struct {
	DefaultProperties
	CompanyID uint64
	Company   Company
	Name      string `gorm:"size:250"`
	Status    string `gorm:"size:25"`
}
type TerritoryArea struct {
	DefaultProperties
	CompanyID  uint64
	Company    Company
	DivisionID uint64
	Division   Division
	AreaID     uint64
	Area       Area
	Name       string `gorm:"size:250"`
	Status     string `gorm:"size:25"`
}
type FiscalYear struct {
	DefaultProperties
	CompanyID uint64
	Company   Company
	Year      int
	Status    string    `gorm:"size:25"`
	YearStart time.Time `form:"YearStart" time_format:"2006-01-02"`
	YearEnd   time.Time `form:"YearEnd" time_format:"2006-01-02"`
	Holidays  []Holiday
	ODs       []OD
}

type Shift struct {
	DefaultProperties
	CompanyID     uint64
	Company       Company
	Name          string `gorm:"size:250"`
	Status        string `gorm:"size:25"`
	StartTime     string `gorm:"size:25"`
	EndTime       string `gorm:"size:25"`
	ShiftCapacity int
	ODs           []OD
}
type Module struct {
	DefaultProperties
	CompanyID   uint64
	Company     Company
	Name        string `gorm:"size:250"`
	FaIcon      string `gorm:"size:25"`
	OrderNumber int
	Status      string `gorm:"size:25"`
	Options     []Option
	Menus       []Menu
}
type Menu struct {
	DefaultProperties
	CompanyID   uint64
	Company     Company
	Module      Module
	ModuleID    uint64
	Name        string `gorm:"size:250"`
	Slug        string `gorm:"size:250"`
	OrderNumber int
	Status      string `gorm:"size:25"`
	Options     []Option
}
type Option struct {
	DefaultProperties
	CompanyID uint64
	Company   Company
	Module    Module
	ModuleID  uint64
	Menu      Menu
	MenuID    uint64
	Type      string `gorm:"size:25"`
	Name      string `gorm:"size:250"`
	Status    string `gorm:"size:25"`
}

type Employee struct {
	DefaultProperties
	Attendances []Attendance
	ODs         []OD
	Customers   []Customer
	// appointment information

	UserID          uint64
	User            User
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
	AppointmentType string    `gorm:"size:25"`
	AppointmentDate time.Time `form:"AppointmentDate"  time_format:"2006-01-02"`
	JoiningDate     time.Time `form:"JoiningDate"  time_format:"2006-01-02"`
	EprID           string    `gorm:"size:25"`
	EmployeeID      string    `gorm:"size:25"`
	EmployeeType    string    `gorm:"size:25"`
	Designation     string    `gorm:"size:250"`
	EmployeeImage   string    `gorm:"size:250"`
	Weekend         string    `gorm:"size:250"`

	// employee information
	EmployeeNameBangla     string    `gorm:"size:250"`
	EmployeeNameEnglish    string    `gorm:"size:250"`
	FatherName             string    `gorm:"size:250"`
	MotherName             string    `gorm:"size:250"`
	HusbandOrWifeName      string    `gorm:"size:250"`
	DateOfBirth            time.Time `form:"DateOfBirth"  time_format:"2006-01-02"`
	PresentAddress         string    `gorm:"size:512"`
	PermanentAddress       string    `gorm:"size:512"`
	EducationQualification string    `gorm:"size:512"`
	ExtraQualification     string    `gorm:"size:512"`
	Nationality            string    `gorm:"size:250"`
	Gender                 string    `gorm:"size:25"`
	BloodGroup             string    `gorm:"size:25"`
	MaritalStatus          string    `gorm:"size:25"`
	Religion               string    `gorm:"size:25"`
	Mobile                 string    `gorm:"size:25"`
	Phone                  string    `gorm:"size:25"`
	Email                  string    `gorm:"size:256;not null"`
	Password               string    `gorm:"-"`

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
	PHHeadMapping        string `gorm:"size:512"`
	LoanHeadMapping      string `gorm:"size:512"`
	Status               string `gorm:"size:25"`
}

type Holiday struct {
	DefaultProperties
	CompanyID    uint64
	Company      Company
	BranchID     uint64
	Branch       Branch
	FiscalYearID uint64
	FiscalYear   FiscalYear
	Name         string    `gorm:"size:250"`
	FromDate     time.Time `form:"FromDate"  time_format:"2006-01-02"`
	ToDate       time.Time `form:"ToDate"  time_format:"2006-01-02"`
	HolidayImage string    `gorm:"size:512"`
	Status       string    `gorm:"size:25"`
}
type Attendance struct {
	DefaultProperties

	Employee   Employee
	EmployeeID uint64
	Mobile     string `gorm:"size:25"`
	InTime     string
	OutTime    string
	Remarks    string `gorm:"size:250"`
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
	Remarks      string    `gorm:"size:250"`
	ODTo         time.Time `form:"ODTo" time_format:"2006-01-02"`
	ODFrom       time.Time `form:"ODFrom" time_format:"2006-01-02"`
	InTime       string
	OutTime      string
}

type Customer struct {
	DefaultProperties
	//Customer Details
	Name            string `gorm:"size:250"`
	ShortCode       string `gorm:"size:250"`
	Address         string `gorm:"size:250"`
	ShippingAddress string `gorm:"size:250"`
	//Contact Person Information
	ContactPerson string `gorm:"size:250"`
	Designation   string `gorm:"size:250"`
	Gender        string `gorm:"size:25"`
	Nationality   string `gorm:"size:25"`
	Employee      Employee
	EmployeeID    uint64
	Mobile        string `gorm:"size:25"`
	Phone         string `gorm:"size:25"`
	Fax           string `gorm:"size:25"`
	Email         string `gorm:"size:250"`
}

type Distributor struct {
	DefaultProperties
	Name            string `gorm:"size:250"`
	ShortCode       string `gorm:"size:25"`
	Address         string `gorm:"size:512"`
	BillingAddress  string `gorm:"size:512"`
	ShippingAddress string `gorm:"size:512"`
	//Contact Person Information
	ContactPerson string `gorm:"size:250"`
	Designation   string `gorm:"size:250"`
	Gender        string `gorm:"size:25"`
	Currency      string `gorm:"size:25"`
	Attention     string `gorm:"size:250"`
	Mobile        string `gorm:"size:25"`
	Phone         string `gorm:"size:25"`
	Fax           string `gorm:"size:25"`
	Email         string `gorm:"size:250"`
}
type Importer struct {
	DefaultProperties
	Name            string `gorm:"size:250"`
	ShortCode       string `gorm:"size:25"`
	Address         string `gorm:"size:512"`
	BillingAddress  string `gorm:"size:512"`
	ShippingAddress string `gorm:"size:512"`
	ContactPerson   string `gorm:"size:250"`
	Designation     string `gorm:"size:250"`
	Currency        string `gorm:"size:25"`
	Mobile          string `gorm:"size:25"`
	Phone           string `gorm:"size:25"`
	Fax             string `gorm:"size:25"`
	Email           string `gorm:"size:250"`
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
	Name string `gorm:"size:250"`
}
type Leave struct {
	DefaultProperties
	Name      string `gorm:"size:250"`
	Total     int
	Status    string `gorm:"size:25"`
	Company   Company
	CompanyID uint64
}
