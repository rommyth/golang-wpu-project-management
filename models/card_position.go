package models

import (
	"project-management/models/types"

	"github.com/google/uuid"
)

type CardPosition struct {
	InternalID int64           `json:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID   uuid.UUID       `json:"public_id" gorm:"column:public_id;uuid;not null"`
	ListID     int64           `json:"list_internal_id" gorm:"column:list_internal_id;not null"`
	CardOrder  types.UUIDArray `json:"card_order" gorm:"column:card_order;type:uuid[]"`
}
