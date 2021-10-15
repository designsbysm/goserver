package database

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	AuthToken   string
	RoleID      uint
	Role        Role   `gorm:"constraint:OnDelete:SET NULL;"`
	RawPassword string `gorm:"-"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" && u.RawPassword == "" {
		return errorPasswordRequired
	} else if u.Password != "" {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.RawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) Create() error {
	return DB.FirstOrCreate(&u, u).Error
}

func (u *User) Read(flags int) error {
	db := DB

	if flags&PreloadRole != 0 {
		db = DB.Preload("Role")
	}

	return db.First(&u, u).Error
}

func (u *User) Update() error {
	return DB.Save(&u).Error
}

func (u *User) ValidatePassword() error {
	if u.RawPassword == "" {
		return errorPasswordRequired
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.RawPassword))
}
