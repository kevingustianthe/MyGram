package models

import (
	"MyGram/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Username     string        `gorm:"unique; not null" json:"username" validate:"required"`
	Email        string        `gorm:"unique; not null" json:"email" validate:"required,email"`
	Password     string        `gorm:"not null" json:"password,omitempty" validate:"required,min=6"`
	Age          uint          `gorm:"not null" json:"age,omitempty" validate:"required,gte=8"`
	CreatedAt    *time.Time    `json:"created_at,omitempty"`
	UpdatedAt    *time.Time    `json:"updated_at,omitempty"`
	Photos       []Photo       `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos,omitempty"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias,omitempty"`
	Comments     []Comment     `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	u.Password = utils.HashPass(u.Password)
	err = nil
	return
}
