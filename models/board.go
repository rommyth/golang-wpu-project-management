package models

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	InternalID    int64      `json:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID      uuid.UUID  `json:"public_id" gorm:"column:public_id"`
	Title         string     `json:"title" gorm:"column:title"`
	Description   string     `json:"description" gorm:"column:description"`
	OwnerID       int64      `json:"owner_internal_id" gorm:"column:owner_internal_id"`
	OwnerPublicID uuid.UUID  `json:"owner_public_id" gorm:"column:owner_public_id"`
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at"`
	Duedate       *time.Time `json:"due_date,omitempty" gorm:"column:due_date"`
}
