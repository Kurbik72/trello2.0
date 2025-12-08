package models

import (
	"time"
)

type DefaultBackground struct {
	ID        string    `gorm:"type:text;primary_key" json:"id"`
	Src       string    `gorm:"type:varchar(500);not null" json:"src"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at" json:"createdAt"`

	Boards []Board `gorm:"foreignKey:BackgroundID" json:"boards,omitempty"`
}

func (DefaultBackground) TableName() string {
	return "default_backgrounds"
}

