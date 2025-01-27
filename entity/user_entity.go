package entity

import (
	"github.com/Caknoooo/go-gin-clean-starter/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Instansi   string `json:"instansi"`
	TelpNumber string `json:"telp_number"`
	InfoFrom   string `json:"info_from"`
	Jenjang    string `json:"jenjang"`
	Role       string `json:"role"`
	IsVerified bool   `json:"is_verified"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	// u.ID = uuid.New()
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
