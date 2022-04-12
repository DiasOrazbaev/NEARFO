package models

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	ID    string  `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
