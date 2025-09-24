package repository

import "crud-all-project/entity"

type ClienteRepository interface {
	FindAll() ([]entity.Cliente, error)
	FindByID(id uint) (*entity.Cliente, error)
	Save(c *entity.Cliente) (*entity.Cliente, error)
	Delete(id uint) error
	FindByDocumento(documento string) (*entity.Cliente, error)
}
