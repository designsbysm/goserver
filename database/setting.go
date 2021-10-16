package database

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Setting struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Key       string          `gorm:"uniqueIndex;not null"`
	Value     datatypes.JSONMap
}

func (s *Setting) GetBool(property string) bool {
	value := s.Value[property]
	if value == nil {
		return false
	}

	return value.(bool)
}

func (s *Setting) Read() error {
	if err := DB.First(&s, s).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	if len(s.Value) == 0 {
		s.Value = make(map[string]interface{})
	}

	return nil
}

func (s *Setting) Upsert() error {
	if s.ID != 0 {
		return DB.Save(&s).Error
	}

	value := s.Value

	err := DB.FirstOrCreate(&s, Setting{
		Key: s.Key,
	}).Error
	if err != nil {
		return err
	}

	s.Value = value

	return DB.Save(&s).Error
}
