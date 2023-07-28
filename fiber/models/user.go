package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primarykey,type:uuid" json:"id"`
	Username string    `json:"username" validate:"required"`
	Email    string    `gorm:"unique" json:"email" validate:"required"`
	Password string    `json:"password,omitempty" validate:"required"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
