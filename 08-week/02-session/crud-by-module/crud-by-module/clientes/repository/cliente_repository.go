package repository

import "crud-by-module/clientes/entity"

type ClienteRepository interface {
	FindAll() ([]entity.Cliente, error)
	FindByID(id uint) (*entity.Cliente, error)
	Save(cliente *entity.Cliente) (*entity.Cliente, error)
	Update(cliente *entity.Cliente) (*entity.Cliente, error)
	Delete(id uint) error
}
