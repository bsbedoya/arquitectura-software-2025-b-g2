package service

import "crud-all-project/entity"

type ClienteService interface {
	FindAll() ([]entity.Cliente, error)
	FindByID(id uint) (*entity.Cliente, error)
	Save(c *entity.Cliente) (*entity.Cliente, error)
	Delete(id uint) error
}
