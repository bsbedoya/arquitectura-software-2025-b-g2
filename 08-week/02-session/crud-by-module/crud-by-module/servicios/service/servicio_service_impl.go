package service

import (
	"crud-by-module/servicios/entity"
	"crud-by-module/servicios/repository"
)

type servicioServiceImpl struct {
	repo repository.ServicioRepository
}

func NewServicioService(r repository.ServicioRepository) ServicioService {
	return &servicioServiceImpl{repo: r}
}

func (s *servicioServiceImpl) FindAll() ([]entity.Servicio, error)        { return s.repo.FindAll() }
func (s *servicioServiceImpl) FindByID(id uint) (*entity.Servicio, error) { return s.repo.FindByID(id) }
func (s *servicioServiceImpl) Save(srv *entity.Servicio) (*entity.Servicio, error) {
	return s.repo.Save(srv)
}
func (s *servicioServiceImpl) Delete(id uint) error { return s.repo.Delete(id) }
func (s *servicioServiceImpl) FindByNombre(nombre string) (*entity.Servicio, error) {
	return s.repo.FindByNombre(nombre)
}
