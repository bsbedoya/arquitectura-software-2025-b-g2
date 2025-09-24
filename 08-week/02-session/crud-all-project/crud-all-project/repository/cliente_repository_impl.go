package repository

import (
	"crud-all-project/database"
	"crud-all-project/entity"

	"gorm.io/gorm"
)

type clienteRepositoryImpl struct {
	db *gorm.DB
}

func NewClienteRepository() ClienteRepository {
	return &clienteRepositoryImpl{db: database.DB}
}

func (r *clienteRepositoryImpl) FindAll() ([]entity.Cliente, error) {
	var clientes []entity.Cliente
	if err := r.db.Find(&clientes).Error; err != nil {
		return nil, err
	}
	return clientes, nil
}

func (r *clienteRepositoryImpl) FindByID(id uint) (*entity.Cliente, error) {
	var c entity.Cliente
	if err := r.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *clienteRepositoryImpl) Save(c *entity.Cliente) (*entity.Cliente, error) {
	if err := r.db.Save(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *clienteRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.Cliente{}, id).Error
}

func (r *clienteRepositoryImpl) FindByDocumento(documento string) (*entity.Cliente, error) {
	var c entity.Cliente
	if err := r.db.Where("documento = ?", documento).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
