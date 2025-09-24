package entity

type Servicio struct {
	IDServicio       uint    `json:"idServicio" gorm:"primaryKey"`
	NombreServicio   string  `json:"nombreServicio"`
	Descripcion      string  `json:"descripcion"`
	Precio           float64 `json:"precio"`
	DuracionEstimada string  `json:"duracionEstimada"`
}
