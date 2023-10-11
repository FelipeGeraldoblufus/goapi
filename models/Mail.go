package models

import (
	"time"

	"gorm.io/gorm"
)

type Mail struct {
	gorm.Model

	OrderID uint `gorm:"unique_index"`
	Rut     string
	Correo  string
	Fecha   time.Time
	Detalle []string
	Total   int
}
