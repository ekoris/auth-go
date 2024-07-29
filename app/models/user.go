package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type User struct {
	orm.Model
	ID       string `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	orm.SoftDeletes
}

// Implementasi BeforeCreate untuk menetapkan UUID
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
