package models

import "time"

type Weekend struct {
	DefaultProperties
	Day    string
	UserID uint64 `gorm:"uniqueIndex;not null"`
}

type Employee struct {
	DefaultProperties

	// appointment information
	Company         string
	Branch          string
	Department      string
	Section         string
	AppointmentType string
	Shift           string
	AppointmentDate time.Time `gorm:"type:date;default:null"`
	JoiningDate     time.Time `gorm:"type:date;default:null"`
	EPRID           uint64
	EmployeeID      uint64
	EmployeeType    string
	Designation     string
	Weekend         string
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
