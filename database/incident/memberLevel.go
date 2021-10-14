package incident

import "gorm.io/gorm"

type MemberLevel struct {
	gorm.Model
	Name string
}
