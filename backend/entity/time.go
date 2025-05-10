package entity

import "gorm.io/gorm"

type Times struct {
	gorm.Model
	Time string `json:"time"`
}
