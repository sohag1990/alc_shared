package models

import (
	"time"
)

type Patient struct {
	DefaultProperties
	UserID        uint64
	Name          string `gorm:"size:250"`
	Address       string `gorm:"size:512"`
	Phone         string `gorm:"size:20"`
	Mobile        string `gorm:"size:100"`
	Gender        string `gorm:"size:25"`
	FatherName    string
	MotherName    string
	DateOfBirth   time.Time
	Occupation    string
	MaritalStatus string
	DateOfMarry   time.Time
	CaseFile      CaseFile
	CaseFileID    uint64

	ProfilePicture string

	Deseases []Desease
	Doctor   Doctor
	DoctorID uint64
}
type CaseFile struct {
	DefaultProperties
	SelectedRemedys  []SelectedRemedy
	PatientID        uint64
	PatientAudio     string
	PatientVideo     string // important for future use of AI, when AI can analyse video and suggest property medicine
	PatientAudioText string // converted from the audio
	DoctorComments   []DoctorComment
	Prescriptions    []Prescription
	PatientComments  []PatientComment
	// KENT'S & Boericke REPERTORY

	// common kent and boericke
	Mind             string
	Vertigo          string
	Head             string
	Eye              string
	Vision           string
	Ear              string
	Hearing          string
	Nose             string
	Face             string
	Mouth            string
	Teeth            string
	Throat           string
	ExternalThroat   string
	Stomach          string
	Abdomen          string
	Rectum           string
	Stool            string
	UrinaryOrgans    string
	Bladder          string
	Kidneys          string
	ProstateGland    string
	Urethra          string
	Urine            string
	GenitaliaMale    string
	GenitaliaFemale  string
	LarynxAndTrachea string
	Respiration      string
	Cough            string
	Expectoration    string
	Chest            string
	Back             string
	Extremities      string
	Sleep            string
	Chill            string
	Fever            string
	Perspiration     string
	Skin             string
	Generalities     string

	// uncommon boericke
	Tongue            string
	Taste             string
	Gums              string
	CirculatorySystem string
	LocomotorSystem   string
	RespiratorySystem string
	NervousSystem     string
	Modalities        string

	/// need improvement of which side of body??? nich theke upore ??? upor theke niche ??? dudh posondo ?? tok na jhal
}

type SelectedRemedy struct {
	DefaultProperties
	Name       string
	RemedyType string
	Comment    string
	CaseFileID uint64
}

type DoctorComment struct {
	DefaultProperties
	CommentType string // 1) General comment to patient 2) Diagonstic comment on prescription 3) improvement comment
	Comment     string
	CaseFileID  string
	Visibility  string // 1) hide 2) show 3) linethrough for metigate the comment with new comment
}
type PatientComment struct {
	DefaultProperties
	CommentType string // 1) General comment to patient 2) Diagonstic comment on prescription 3) improvement comment
	Comment     string
	CaseFileID  string
	Visibility  string // 1) hide 2) show 3) linethrough for metigate the comment with new comment
}
type Prescription struct {
	DefaultProperties
	CaseFileID uint64
	Remedy     string
	Dosage     string
	Comment    string
	Visibility string // 1) hide 2) show 3) linethrough for metigate the comment with new medicine
}

// for software data entry
// type

// // for data entry
// type Symptom struct {
// 	DefaultProperties
// 	Title        string
// 	AffectedArea []AffectedArea `gorm:"many2many:symptom_affectedAreas;"`
// }

// for data entry
//
//	type AffectedArea struct {
//		DefaultProperties
//		Title string
//	}
type Doctor struct {
	DefaultProperties
	UserID          uint64
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
	Name        string
	OMIMID      string
	Status      bool
	Anatomicals []Anatomical `gorm:"many2many:desease_anatomicals;"`
	Globals     []Global     `gorm:"many2many:desease_globals;"`
	Picture     string
}
type Anatomical struct {
	DefaultProperties
	Category string
	Count    int
}
type Global struct {
	DefaultProperties
	Category string
	Count    int
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

/// software database for pre data entry

type remedy struct {
	ID              int
	SymptomCategory string
	Symptom         string
	Description     string
	Remedy          string
	Antitude        string
	RelatedRemedy   string
	// need AI help
	// 	Solitude

	// Aversion to -
	// Desire for
	// 	Duration of time changed

	// Passes too rapidly
	// Passes too slowly
}

// Entities:

// Symptom: Represents a specific symptom or mental condition.

// Attributes:
// Symptom_ID (Primary Key)
// Symptom_Name
// Description
// Remedy: Represents a remedy or treatment associated with a symptom.

// Attributes:
// Remedy_ID (Primary Key)
// Remedy_Name
// Symptom_ID (Foreign Key referencing Symptom)
// Remedy_Category: Represents categories or types of remedies.

// Attributes:
// Category_ID (Primary Key)
// Category_Name
// Remedy_RemedyCategory: Many-to-Many relationship between Remedy and Remedy_Category.

// Attributes:
// Remedy_ID (Foreign Key referencing Remedy)
// Category_ID (Foreign Key referencing Remedy_Category)
// Remedy_Symptom: Many-to-Many relationship between Remedy and Symptom (to represent multiple remedies for a symptom).

// Attributes:
// Remedy_ID (Foreign Key referencing Remedy)
// Symptom_ID (Foreign Key referencing Symptom)
