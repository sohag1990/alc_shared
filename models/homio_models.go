package models

import (
	"time"
)

type Patient struct {
	DefaultProperties
	Name                string `gorm:"size:250"`
	Address             string `gorm:"size:512"`
	Phone               string `gorm:"size:20"`
	Mobile              string `gorm:"size:100"`
	Gender              string `gorm:"size:25"`
	PatientHistory      string
	PatientHistoryAudio string
	PatientHistoryVideo string
	ProfilePicture      string

	Deseases []Desease
	Doctor   Doctor
	DoctorID uint64
}
type Doctor struct {
	DefaultProperties
	Name            string `gorm:"size:250"`
	Degree          string `gorm:"size:250"`
	Address         string `gorm:"size:250"`
	Mobile          string `gorm:"size:250"`
	Email           string `gorm:"size:250"`
	Bio             string `gorm:"size:250"`
	DateOfBirth     time.Time
	Picture         string
	BankAccount     string `gorm:"size:250"`
	BankName        string `gorm:"size:250"`
	Chamber         string `gorm:"size:250"`
	Speciality      string `gorm:"size:250"`
	Nationality     string `gorm:"size:250"`
	Patients        []Patient
	Fee             int
	ChamberFee      int
	DutyHour        string `gorm:"size:250"`
	OnlineStatus    bool
	Verification    bool
	Status          bool
	TermsConditions string // if any doctor try to be direct of paitent that will be violance
}

type PatientPaymentHistory struct {
	DefaultProperties
	Fee               int
	DoctorID          uint64
	PatientID         uint64
	PaymentDate       time.Time
	PaymentMethod     string // Bkash, Nagod, Rocket, UCash
	PaymentAccount    string
	TransactionNumber string
	Status            string
}
type DoctorVerification struct {
	DefaultProperties
	Name     string
	Type     string //physical verificate or online verification
	Files    string
	Verified string
}
type Desease struct {
	DefaultProperties
	Name              string
	Type              string //sextual, Birth, Deaf, Blindness, Anemia etc
	Temperature       int
	BloodPressure     int
	Date              time.Time
	DoctorID          uint64
	Doctor            Doctor
	PatientID         uint64
	Patient           Patient
	Status            bool
	MedicineHistories []MedicineHistory
	Picture           string
}

type MedicineHistory struct {
	DefaultProperties
	Name        string
	Potency     int
	Date        time.Time
	SuccessRate int // out of 100
	DeseaseID   uint64
	Desease     Desease
}
type Medicine struct {
	DefaultProperties
	Name    string
	Potency int
	Made    string
}
