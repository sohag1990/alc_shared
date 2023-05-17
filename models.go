package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// @DefaultProperties
// DefaultProperties Auto gen ID Primary, CreatedAt, UpdatedAt, DeletedAt Property
type DefaultProperties struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time            `json:"-"`
	DeletedAt *time.Time            `sql:"index" json:"-"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
