package incident

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	CallSign string
	Members  []Member `gorm:"many2many:unit_members;"`
}
