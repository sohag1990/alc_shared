package models

import "time"

type HRAppointment struct {
	DefaultProperties
	Company         string
	Branch          string
	Department      string
	Section         string
	AppointmentType string
	Shift           string
	AppointmentDate time.Time
	JoiningDate     time.Time
	EPRID           uint64
	EmployeeID      uint64
	EmployeeType    string
	Designation     string
	Weekend         string
	EmployeeImage   string
}

type Weekend struct {
	DefaultProperties
	Day    string
	UserID uint64 `gorm:"uniqueIndex;not null"`
}

type Employee struct {
	DefaultProperties
	EmployeeNameBangla      string
	EmployeeNameEnglish     string
	FatherName              string
	MotherName              string
	HusbandOrWifeName       string
	DateOfBirth             time.Time
	PresentAddressBangla    string
	PresentAddressEnglish   string
	PermanentAddressBangla  string
	PermanentAddressEnglish string
	EducationQualification  string
	Nationality             string
	Gender                  string
	BloodGroup              string
	MaritalStatus           string
	Religion                string
	Mobile                  string
	Phone                   string
	Email                   string
}

type SalaryStructure struct {
	DefaultProperties
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
