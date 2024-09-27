package models

import (
	"time"

	"gorm.io/gorm"
)

type FrontendUserData struct {
	DefaultProperties
	FullName     string
	Email        string
	Role         string
	Hostname     string
	HostnameOnly string
	ButtonText   string
	Url          string
}

type User struct {
	DefaultProperties
	Email     string `gorm:"unique;size:256;not null"`
	Password  string `form:"password" json:"password" json:"-"`
	Role      string
	Verified  bool `json:"verified"`
	Hostname  string
	ProfileID uint64
	Profile   Profile
	Property  []Property
	Api       string
}

type DefaultProperties struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time `json:"-"`
	Deleted   gorm.DeletedAt
}

type EmailSetting struct {
	DefaultProperties
	FromEmail      string
	CcEmail        string
	BccEmail       string
	AwsZone        string
	AWSAccessKeyID string
	AWSSecretKeyID string
	EmailSend      bool
}

type Leads struct {
	DefaultProperties
	Niche    string `gorm:"size:30"`
	Other    string `gorm:"size:350"`
	Cuntry   string `gorm:"size:20"`
	City     string `gorm:"size:50"`
	Name     string `gorm:"size:250"`
	Email    string `gorm:"unique;size:256;not null"`
	Phone    string `gorm:"size:15"`
	Address  string `gorm:"size:250"`
	Company  string `gorm:"size:250"`
	Website  string `gorm:"size:250"`
	Position string `gorm:"size:250"`
	Status   string `gorm:"size:10"`
	Note     string `gorm:"size:512"`
}
