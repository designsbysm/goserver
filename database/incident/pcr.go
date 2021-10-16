package incident

import (
	"time"

	"gorm.io/gorm"
)

type PCR struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	DeletedAt  *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	IncidentID uint
	PatientID  uint
	Patient    Patient
	UnitID     uint
	Unit       Unit
	AuthorID   uint
	Author     Member
}
