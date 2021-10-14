package incident

import "gorm.io/gorm"

type PCR struct {
	gorm.Model
	IncidentID uint
	PatientID  uint
	Patient    Patient
	UnitID     uint
	Unit       Unit
	AuthorID   uint
	Author     Member
}
