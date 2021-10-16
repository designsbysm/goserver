package incident

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	CallSign  string
	Members   []Member `gorm:"many2many:unit_members;"`
}
