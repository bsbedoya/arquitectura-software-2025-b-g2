package dto

type VehiculoDTO struct {
	Placa     string `json:"placa"`
	Marca     string `json:"marca"`
	Modelo    string `json:"modelo"`
	Tipo      string `json:"tipo"`
	ClienteID uint   `json:"cliente_id"`
}
