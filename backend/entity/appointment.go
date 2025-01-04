package entity

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model

	Date    time.Time `json:"date"`
	Illness string    `json:"illness"`

	DepartmentID uint
	Department   Departments `json:"department" gorm:"foreignKey:DepartmentID"`

	TimeID uint
	Time   Times `json:"time" gorm:"foreignKey:TimeID"`

	RoomID uint
	Room   Room `json:"room" gorm:"foreignKey:RoomID"`
}
