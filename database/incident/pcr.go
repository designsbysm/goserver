package incident

import (
	"time"

	"gorm.io/gorm"
)

type InPCR struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	DeletedAt  *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	IncidentID uint
	PatientID  uint
	Patient    InPatient
	UnitID     uint
	Unit       InUnit
	AuthorID   uint
	Author     InMember
}
