package repository

import "crud-all-project/entity"

type ServicioRepository interface {
	FindAll() ([]entity.Servicio, error)
	FindByID(id uint) (*entity.Servicio, error)
	Save(s *entity.Servicio) (*entity.Servicio, error)
	Delete(id uint) error
	FindByNombre(nombre string) (*entity.Servicio, error)
}
