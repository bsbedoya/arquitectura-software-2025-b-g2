package dto

type ClienteDTO struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}
