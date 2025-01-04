package entity

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model

	Date    time.Time `json:"date"`
	Illness string    `json:"illness"`

	TimeID uint
	Time   Times `json:"time" gorm:"foreignKey:TimeID"`

	DepartmentID uint
	Department   Departments `json:"department" gorm:"foreignKey:DepartmentID"`

	PatientID *uint
	Patient   Patient `json:"patient" gorm:"foreignKey:PatientID"`
}
