package dto

type ClienteDTO struct {
	NombreCompleto string `json:"nombreCompleto"`
	Documento      string `json:"documento"`
	Telefono       string `json:"telefono"`
	Correo         string `json:"correo"`
}
