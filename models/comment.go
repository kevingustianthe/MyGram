package models

import "time"

type Comment struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	PhotoID   uint       `gorm:"not null" json:"photo_id"`
	Message   string     `gorm:"not null" json:"message" validate:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *User      `json:",omitempty"`
	Photo     *Photo     `json:",omitempty"`
}
