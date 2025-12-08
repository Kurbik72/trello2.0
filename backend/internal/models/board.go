package models

import (
	"time"
)

type Board struct {
	ID            string     `gorm:"type:text;primary_key" json:"id"`
	Title         string     `gorm:"type:varchar(255);not null" json:"title"`
	IsFavorite    bool       `gorm:"default:false;column:is_favorite" json:"isFavorite"`
	LinkToBoard   string     `gorm:"type:varchar(500);not null;column:link_to_board" json:"linkToBoard"`
	BackgroundSrc string     `gorm:"type:varchar(500);not null;column:background_src" json:"backgroundSrc"`
	UserID        string     `gorm:"type:text;not null;column:user_id;index" json:"userId"`
	BackgroundID  *string    `gorm:"type:text;column:background_id;index" json:"backgroundId,omitempty"`
	CreatedAt     time.Time  `gorm:"autoCreateTime;column:created_at" json:"createdAt"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`

	User      User              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Background *DefaultBackground `gorm:"foreignKey:BackgroundID;constraint:OnDelete:SET NULL" json:"background,omitempty"`
}

func (Board) TableName() string {
	return "boards"
}

