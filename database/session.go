package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex" json:"id"`
	FirstName string    `gorm:"-" json:"firstName"`
	LastName  string    `gorm:"-" json:"lastName"`
	Role      string    `gorm:"-" json:"role"`
	Token     string    `json:"token,omitempty"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()

	return nil
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
