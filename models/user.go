package models

import (
	"html/template"
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

type Lead struct {
	DefaultProperties
	Niche         string `gorm:"size:30"`
	Other         string `gorm:"size:350"`
	Country       string `gorm:"size:20"`
	City          string `gorm:"size:50"`
	Name          string `gorm:"size:250"`
	Email         string `gorm:"unique;size:256;not null"`
	Phone         string `gorm:"size:15"`
	Address       string `gorm:"size:250"`
	Company       string `gorm:"size:250"`
	Website       string `gorm:"size:250"`
	Position      string `gorm:"size:250"`
	Instagram     string `gorm:"size:512"`
	Facebook      string `gorm:"size:512"`
	Youtube       string `gorm:"size:512"`
	X             string `gorm:"size:512"`
	Linkedin      string `gorm:"size:512"`
	Placeid       string `gorm:"size:100"`
	Cid           string `gorm:"size:100"`
	Category      string `gorm:"size:512"`
	ReviewCount   string `gorm:"size:10"`
	AverageRating string `gorm:"size:512"`
	Latitude      string `gorm:"size:512"`
	Longitude     string `gorm:"size:512"`
}

type Campaign struct {
	DefaultProperties
	Name            string `gorm:"size:250"`
	Type            string `gorm:"size:250"` //"Email, SMS"
	Status          string `gorm:"size:25"`
	SendTime        *time.Time
	UserID          uint64
	User            User
	CampaignLeads   []CampaignLead
	CampaignContent CampaignContent
}

type CampaignLead struct {
	DefaultProperties
	CampaignID  uint64
	LeadID      uint64
	Unsubscribe bool
	SentAt      *time.Time
	OpendAt     *time.Time
	ClickedAt   *time.Time
	Status      bool
}
type CampaignContent struct {
	DefaultProperties
	CampaignID uint64
	UserID     uint64
	Subject    string `gorm:"size:512"`
	Body       template.HTML
}

type CampaignTemplate struct {
	DefaultProperties
	UserID  uint64
	Type    string `gorm:"size:25"`
	Name    string `gorm:"size:125"`
	Subject string `gorm:"size:512"`
	Body    template.HTML
}

type CampaignAnalytics struct {
	DefaultProperties
	UserID       uint64
	CampaignID   uint64
	TotalSent    int
	TotalOpened  int
	TotalClicked int
	TotalFailed  int
	Score        string `gorm:"size:25"`
}
