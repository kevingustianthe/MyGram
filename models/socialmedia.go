package models

import "time"

type SocialMedia struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"not null" json:"name" validate:"required"`
	SocialMediaURL string     `gorm:"not null" json:"social_media_url" validate:"required"`
	UserID         uint       `gorm:"not null" json:"user_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	User           *User      `json:",omitempty"`
}
