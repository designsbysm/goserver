package database

import (
	"time"
)

type Session struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserID    uint      `gorm:"uniqueIndex" json:"id"`
	FirstName string    `gorm:"-" json:"firstName"`
	LastName  string    `gorm:"-" json:"lastName"`
	Role      string    `gorm:"-" json:"role"`
	Token     string    `json:"token"`
}

func (s *Session) Read() error {
	return DB.First(&s, s).Error
}

func (s *Session) Upsert() error {
	tempToken := s.Token

	if err := DB.FirstOrCreate(&s, Session{
		UserID: s.UserID,
	}).Error; err != nil {
		return err
	}
	s.Token = tempToken

	return DB.Save(&s).Error
}
