package entity

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name         string `json:"name"`
	DepartmentID uint   `json:"department_id"` // Foreign Key สำหรับเชื่อมโยงกับแผนก

	// การเชื่อมโยงไปยัง Department
	Department Departments `gorm:"foreignKey:DepartmentID" json:"department"`
}
