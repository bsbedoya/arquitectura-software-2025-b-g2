package entity

type Vehiculo struct {
	IDVehiculo   uint   `gorm:"primaryKey;autoIncrement" json:"idVehiculo"`
	Placa        string `gorm:"index" json:"placa"`
	TipoVehiculo string `json:"tipoVehiculo"`
	IDCliente    uint   `json:"idCliente"`
}
