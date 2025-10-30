package models

import "github.com/google/uuid"

type Label struct {
	InternalID int64     `json:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" gorm:"column:public_id"`
	Name       string    `json:"name" gorm:"column:name"`
	Color      string    `json:"color" gorm:"column:color"`
}
