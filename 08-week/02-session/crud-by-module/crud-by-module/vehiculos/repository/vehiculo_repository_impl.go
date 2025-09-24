package repository

import (
	"crud-by-module/vehiculos/entity"

	"gorm.io/gorm"
)

type vehiculoRepositoryImpl struct {
	db *gorm.DB
}

func NewVehiculoRepository(db *gorm.DB) VehiculoRepository {
	return &vehiculoRepositoryImpl{db: db}
}

func (r *vehiculoRepositoryImpl) FindAll() ([]entity.Vehiculo, error) {
	var vehiculos []entity.Vehiculo
	err := r.db.Find(&vehiculos).Error
	return vehiculos, err
}

func (r *vehiculoRepositoryImpl) FindByID(id uint) (*entity.Vehiculo, error) {
	var vehiculo entity.Vehiculo
	err := r.db.First(&vehiculo, id).Error
	return &vehiculo, err
}

func (r *vehiculoRepositoryImpl) Save(v *entity.Vehiculo) (*entity.Vehiculo, error) {
	err := r.db.Create(v).Error
	return v, err
}

func (r *vehiculoRepositoryImpl) Update(v *entity.Vehiculo) (*entity.Vehiculo, error) {
	err := r.db.Save(v).Error
	return v, err
}

func (r *vehiculoRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.Vehiculo{}, id).Error
}

func (r *vehiculoRepositoryImpl) FindByPlaca(placa string) (*entity.Vehiculo, error) {
	var vehiculo entity.Vehiculo
	err := r.db.Where("placa = ?", placa).First(&vehiculo).Error
	return &vehiculo, err
}

func (r *vehiculoRepositoryImpl) FindByTipo(tipo string) ([]entity.Vehiculo, error) {
	var vehiculos []entity.Vehiculo
	err := r.db.Where("tipo = ?", tipo).Find(&vehiculos).Error
	return vehiculos, err
}

func (r *vehiculoRepositoryImpl) FindByIdCliente(idCliente uint) ([]entity.Vehiculo, error) {
	var vehiculos []entity.Vehiculo
	err := r.db.Where("cliente_id = ?", idCliente).Find(&vehiculos).Error
	return vehiculos, err
}
