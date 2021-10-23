package incident

import (
	"time"

	"gorm.io/gorm"
)

type InUnit struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	CallSign  string
	Members   []InMember `gorm:"many2many:in_unit_members;"`
}
