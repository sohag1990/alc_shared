package models

import (
	"time"

	"github.com/go-gormsoft/soft_delete"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string `gorm:"uniqueIndex;not null;size:256"`
	Password        string `form:"password" json:"password" binding:"required" json:"-"`
	Role            string
	Verified        bool `json:"verified"`
	Hostname        string
	Profile         Profile
	Appointment     Appointment
	Weekend         Weekend
	Employee        Employee
	SalaryStructure ServiceCategory
}

type FrontendUserData struct {
	DefaultProperties
	FullName   string
	Email      string
	Role       string
	Hostname   string
	ButtonText string
	Url        string
}

type Appointment struct {
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
	UserID uint64
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

type DefaultProperties struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time            `json:"-"`
	DeletedAt *time.Time            `sql:"index" json:"-"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
