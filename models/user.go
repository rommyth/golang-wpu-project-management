package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	InternalID int64          `json:"internal_id" gorm:"column:internal_id;primaryKey"`
	PublicID   uuid.UUID      `json:"public_id" gorm:"column:public_id"`
	Name       string         `json:"name" gorm:"column:name"`
	Email      string         `json:"email" gorm:"column:email;unique"`
	Password   string         `json:"password" gorm:"column:password"`
	Role       string         `json:"role" gorm:"column:role"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
