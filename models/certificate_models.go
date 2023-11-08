package models

type Fra struct {
	DefaultProperties
	Name    string `gorm:"size:250"`
	Address string `gorm:"size:512"`
	Phone   string `gorm:"size:20"`
}
