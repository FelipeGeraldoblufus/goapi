package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string `gorm:"not null;unique_index"`
	price  uint   `gorm:"not null"` //entero sin signo
	UserID uint
}
