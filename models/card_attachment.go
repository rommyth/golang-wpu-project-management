package models

import (
	"time"

	"github.com/google/uuid"
)

type CardAttachment struct {
	InternalID int64     `json:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" gorm:"column:public_id"`
	CardID     int64     `json:"card_internal_id" gorm:"column:card_internal_id"`
	UserID     int64     `json:"user_internal_id" gorm:"column:user_internal_id"`
	File       string    `json:"file" gorm:"column:file"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}
