package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room         string `json:"room"`
	DepartmentID uint   `json:"department_id"` // Foreign Key สำหรับเชื่อมโยงกับแผนก

	// การเชื่อมโยงไปยัง Department
	Department Departments `gorm:"foreignKey:DepartmentID" json:"department"`
}
