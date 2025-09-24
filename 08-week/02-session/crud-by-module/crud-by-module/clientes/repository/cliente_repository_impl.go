package repository

import (
	"crud-by-module/clientes/entity"
	"crud-by-module/database"
)

type clienteRepositoryImpl struct{}

func NewClienteRepository() ClienteRepository {
	return &clienteRepositoryImpl{}
}

func (r *clienteRepositoryImpl) FindAll() ([]entity.Cliente, error) {
	var clientes []entity.Cliente
	result := database.DB.Find(&clientes)
	return clientes, result.Error
}

func (r *clienteRepositoryImpl) FindByID(id uint) (*entity.Cliente, error) {
	var cliente entity.Cliente
	result := database.DB.First(&cliente, id)
	return &cliente, result.Error
}

func (r *clienteRepositoryImpl) Save(cliente *entity.Cliente) (*entity.Cliente, error) {
	result := database.DB.Create(cliente)
	return cliente, result.Error
}

func (r *clienteRepositoryImpl) Update(cliente *entity.Cliente) (*entity.Cliente, error) {
	result := database.DB.Save(cliente)
	return cliente, result.Error
}

func (r *clienteRepositoryImpl) Delete(id uint) error {
	result := database.DB.Delete(&entity.Cliente{}, id)
	return result.Error
}
