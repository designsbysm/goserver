package incident

import (
	"time"

	"gorm.io/gorm"
)

type Incident struct {
	gorm.Model
	DateOfService  time.Time
	IncidentNumber string
	PCRS           []PCR `gorm:"many2many:incident_pcrs;"`
}
