package incident

import (
	"time"

	"gorm.io/gorm"
)

type InMember struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	FirstName string
	LastName  string
	LevelID   uint
	Level     InMemberLevel
}
