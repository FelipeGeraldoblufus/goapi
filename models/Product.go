package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string `gorm:"not null;unique_index" json:"name"`
	Price  uint   `gorm:"not null" json:"price"` //entero sin signo
	UserID uint   `json:"user_id"`
}
