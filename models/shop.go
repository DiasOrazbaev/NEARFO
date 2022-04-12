package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ID            string  `gorm:"primaryKey"`
	Name          string  `json:"name"`
	Postcode      string  `json:"postcode"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	PostalAddress string  `json:"postal_address"`
}
