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
	ID          uint64 `gorm:"primary_key"`
	EncryptedID string `gorm:"-"` //ignore in crude operation
	CreatedAt   *time.Time
	UpdatedAt   *time.Time `json:"-"`
	Deleted     gorm.DeletedAt
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
	Keyword       string `gorm:"size:150"`
	City          string `gorm:"size:100"`
	Country       string `gorm:"size:50"`
	Name          string `gorm:"size:250"`
	ReviewCount   string `gorm:"size:10"`
	AverageRating string `gorm:"size:512"`
	PriceLevel    string `gorm:"size:50"`
	MainCategory  string `gorm:"size:50"`
	Niche         string `gorm:"size:30"`

	Phones           string `gorm:"size:100"`
	Email            string `gorm:"size:100"`
	Address          string `gorm:"size:250"`
	Website          string `gorm:"size:250"`
	Company          string `gorm:"size:250"`
	MapUrl           string `gorm:"size:512"`
	Photos           string
	Position         string `gorm:"size:250"`
	SocialMediaLinks string `gorm:"size:1200"`
	Placeid          string `gorm:"size:100"`
	Cid              string `gorm:"size:100"`
	Category         string `gorm:"size:512"`

	Latitude  float64
	Longitude float64
	Timestamp time.Time
	Crawled   bool
	ValidSsl  bool
	Title     string `gorm:"size:512"`
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
