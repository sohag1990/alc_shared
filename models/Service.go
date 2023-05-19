package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// Service public model generated by bindu
type Service struct {
	ID                int               `gorm:"primary_key;auto_increment;not_null" json:"ID"`
	ServiceCategories []ServiceCategory `gorm:"many2many:service_servicecategories;"`
	Title             string
	Price             float32
	OfferPrice        float32
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time            `gorm:"index"`
	IsDel             soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	Status            bool
	IsSelected        string
}

// ServiceCategory public model generated by bindu
type ServiceCategory struct {
	ServiceID int `json:"-"`
	Title     string
	ID        int `gorm:"primarykey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time            `gorm:"index"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	Services  []Service             `gorm:"many2many:service_servicecategories;"`
	Checked   string
}