package models

import (
	"time"
)

type User struct {
	ID           string    `gorm:"type:text;primary_key" json:"id"`
	Email        string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null;column:password_hash" json:"-"`
	RefreshToken *string   `gorm:"type:text;column:refresh_token" json:"-"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`

	Boards []Board `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"boards,omitempty"`
}

func (User) TableName() string {
	return "users"
}

