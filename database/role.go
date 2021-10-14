package database

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex;not null"`
	IsAdmin bool
}
