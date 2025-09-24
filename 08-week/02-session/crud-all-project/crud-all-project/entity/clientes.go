package entity

type Cliente struct {
	IDCliente      uint   `json:"idCliente" gorm:"primaryKey"`
	NombreCompleto string `json:"nombreCompleto"`
	Documento      string `json:"documento"`
	Telefono       string `json:"telefono"`
	Correo         string `json:"correo"`
}
