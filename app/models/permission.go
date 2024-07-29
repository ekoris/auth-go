package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type Permission struct {
	orm.Model
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	orm.SoftDeletes
}

// Implementasi BeforeCreate untuk menetapkan UUID
func (permission *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	permission.ID = uuid.New().String()
	return
}
