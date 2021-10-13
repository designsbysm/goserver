package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var errorPasswordRequired = errors.New("User error: password is required")

type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	Password     string `gorm:"not null"`
	PasswordHash string `gorm:"-"`
	AuthToken    string
	RoleID       uint
	Role         Role `gorm:"constraint:OnDelete:SET NULL;"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.PasswordHash != "" {
		u.Password = u.PasswordHash
		return nil
	} else if u.Password == "" {
		return errorPasswordRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) ValidatePassword(password string) error {
	if password == "" {
		return errorPasswordRequired
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
