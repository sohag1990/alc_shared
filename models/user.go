package models

import (
	"time"

	"gorm.io/gorm"
)

type FrontendUserData struct {
	DefaultProperties
	FullName   string
	Email      string
	Role       string
	Hostname   string
	ButtonText string
	Url        string
}

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null;size:256"`
	Password  string `form:"password" json:"password" binding:"required" json:"-"`
	Role      string
	Verified  bool `json:"verified"`
	Hostname  string
	ProfileID uint64 `gorm:"uniqueIndex;not null"`

	// HRM fields
	AppointmentID     uint64 `gorm:"uniqueIndex;not null"`
	EmployeeID        uint64 `gorm:"uniqueIndex;not null"`
	SalaryStructureID uint64 `gorm:"uniqueIndex;not null"`
	Profile           Profile
	Appointment       Appointment
	Weekend           Weekend
	Employee          Employee
	SalaryStructure   ServiceCategory
}

type DefaultProperties struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time `json:"-"`
	Deleted   gorm.DeletedAt
}
