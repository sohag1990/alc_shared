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
