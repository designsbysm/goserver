package incident

import (
	"time"

	"gorm.io/gorm"
)

type Incident struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	DateOfService  time.Time
	IncidentNumber string
	PCRS           []PCR `gorm:"many2many:incident_pcrs;"`
}
