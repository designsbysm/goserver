package incident

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	FirstName string
	LastName  string
	LevelID   uint
	Level     MemberLevel
}
