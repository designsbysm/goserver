package database

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Name      string          `gorm:"uniqueIndex;not null" json:"name"`
	IsAdmin   bool            `json:"isAdmin"`
}

func (r *Role) Create() error {
	return DB.FirstOrCreate(&r, r).Error
}
