package service

import (
	"crud-all-project/entity"
	"crud-all-project/repository"
)

type clienteServiceImpl struct {
	repo repository.ClienteRepository
}

func NewClienteService(r repository.ClienteRepository) ClienteService {
	return &clienteServiceImpl{repo: r}
}

func (s *clienteServiceImpl) FindAll() ([]entity.Cliente, error) {
	return s.repo.FindAll()
}

func (s *clienteServiceImpl) FindByID(id uint) (*entity.Cliente, error) {
	return s.repo.FindByID(id)
}

func (s *clienteServiceImpl) Save(c *entity.Cliente) (*entity.Cliente, error) {
	return s.repo.Save(c)
}

func (s *clienteServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
