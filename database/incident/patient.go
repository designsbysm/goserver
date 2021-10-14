package incident

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}
