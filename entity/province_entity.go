package entity

import "github.com/google/uuid"

type Province struct {
	ID    uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Value string    `json:"value"`
}
