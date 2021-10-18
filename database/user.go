package database

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	FirstName   string          `gorm:"not null" json:"firstName"`
	LastName    string          `gorm:"not null" json:"lastName"`
	Email       string          `gorm:"uniqueIndex;not null" json:"email"`
	Password    string          `gorm:"not null" json:"-"`
	RoleID      uint            `json:"roleID"`
	Role        *Role           `gorm:"constraint:OnDelete:SET NULL;" json:"role,omitempty"`
	Session     Session
	RawPassword string `gorm:"-" json:"password,omitempty"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" && u.RawPassword == "" {
		return ErrPasswordRequired
	} else if u.RawPassword == "" {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.RawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	u.RawPassword = ""

	return nil
}

func (u *User) Create() error {
	return DB.FirstOrCreate(&u, u).Error
}

func (u *User) Delete() error {
	return DB.Delete(&u, u).Error
}

func (u *User) List(flags int) ([]User, error) {
	db := DB
	list := []User{}

	if flags&PreloadRole != 0 {
		db = DB.Preload("Role")
	}

	err := db.Find(&list, u).Error

	return list, err
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

func (u *User) ValidatePassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
