package models

import "time"

type Photo struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"not null" json:"title" validate:"required"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `gorm:"not null" json:"photo_url" validate:"required"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Comments  []Comment  `gorm:"foreignKey:PhotoID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
	User      *User      `json:",omitempty"`
}
