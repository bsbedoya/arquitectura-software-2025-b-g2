package repository

import "crud-all-project/entity"

type VehiculoRepository interface {
	FindAll() ([]entity.Vehiculo, error)
	FindByID(id uint) (*entity.Vehiculo, error)
	Save(v *entity.Vehiculo) (*entity.Vehiculo, error)
	Update(v *entity.Vehiculo) (*entity.Vehiculo, error)
	Delete(id uint) error
	FindByPlaca(placa string) (*entity.Vehiculo, error)
	FindByTipo(tipo string) ([]entity.Vehiculo, error)
	FindByIdCliente(idCliente uint) ([]entity.Vehiculo, error)
}
