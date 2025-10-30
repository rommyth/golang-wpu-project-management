package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	InternalID int64     `json:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" gorm:"column:public_id"`
	CardID     int64     `json:"card_internal_id" gorm:"column:card_internal_id"`
	CardPubID  uuid.UUID `json:"card_id" gorm:"column:card_id"`
	UserID     int64     `json:"user_internal_id" gorm:"column:user_internal_id"`
	UserPubID  uuid.UUID `json:"user_id" gorm:"column:user_id"`
	Message    string    `json:"message" gorm:"column:message"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}
