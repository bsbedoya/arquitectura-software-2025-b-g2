package entity

import "gorm.io/gorm"

type Servicio struct {
	gorm.Model
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Descripcion string  `json:"descripcion"`
}
