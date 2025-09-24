package entity

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}
