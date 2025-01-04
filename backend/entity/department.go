package entity

import "gorm.io/gorm"

type Departments struct {
	gorm.Model
	Department  string `json:"department"`
	Description string `json:"description"`
	HeadOfDept  string `json:"head_of_department"`
}
