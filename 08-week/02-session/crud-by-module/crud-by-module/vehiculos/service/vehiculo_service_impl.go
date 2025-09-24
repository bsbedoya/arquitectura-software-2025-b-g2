package service

import (
	"crud-by-module/vehiculos/entity"
	"crud-by-module/vehiculos/repository"
)

type vehiculoServiceImpl struct {
	repo repository.VehiculoRepository
}

func NewVehiculoService(r repository.VehiculoRepository) VehiculoService {
	return &vehiculoServiceImpl{repo: r}
}

func (s *vehiculoServiceImpl) FindAll() ([]entity.Vehiculo, error) {
	return s.repo.FindAll()
}

func (s *vehiculoServiceImpl) FindByID(id uint) (*entity.Vehiculo, error) {
	return s.repo.FindByID(id)
}

func (s *vehiculoServiceImpl) Save(v *entity.Vehiculo) (*entity.Vehiculo, error) {
	return s.repo.Save(v)
}

func (s *vehiculoServiceImpl) Update(v *entity.Vehiculo) (*entity.Vehiculo, error) {
	return s.repo.Update(v)
}

func (s *vehiculoServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *vehiculoServiceImpl) FindByPlaca(placa string) (*entity.Vehiculo, error) {
	return s.repo.FindByPlaca(placa)
}

func (s *vehiculoServiceImpl) FindByTipo(tipo string) ([]entity.Vehiculo, error) {
	return s.repo.FindByTipo(tipo)
}

func (s *vehiculoServiceImpl) FindByIdCliente(idCliente uint) ([]entity.Vehiculo, error) {
	return s.repo.FindByIdCliente(idCliente)
}
