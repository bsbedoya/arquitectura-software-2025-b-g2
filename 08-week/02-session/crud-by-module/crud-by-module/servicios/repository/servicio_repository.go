package repository

import "crud-by-module/servicios/entity"

type ServicioRepository interface {
	GetAll() ([]entity.Servicio, error)
	GetByID(id int) (entity.Servicio, error)
	Create(servicio entity.Servicio) (entity.Servicio, error)
	Update(servicio entity.Servicio) (entity.Servicio, error)
	Delete(id uint) error
	FindByNombre(nombre string) (*entity.Servicio, error)
	Save(srv *entity.Servicio) (*entity.Servicio, error)
	FindAll() ([]entity.Servicio, error)
	FindByID(id uint) (*entity.Servicio, error)
}
