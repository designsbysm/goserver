package incident

import (
	"time"

	"gorm.io/gorm"
)

type InIncident struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	DateOfService  time.Time
	IncidentNumber string
	PCRS           []InPCR `gorm:"many2many:in_incident_pcrs;"`
}
