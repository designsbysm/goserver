package database

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex;not null"`
	IsAdmin bool
}

func (r *Role) Create() error {
	return DB.FirstOrCreate(&r, r).Error
}
