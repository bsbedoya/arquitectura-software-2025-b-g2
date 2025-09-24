package repository

import (
	"crud-all-project/database"
	"crud-all-project/entity"
)

type vehiculoRepositoryImpl struct{}

func NewVehiculoRepository() VehiculoRepository {
	return &vehiculoRepositoryImpl{}
}

func (r *vehiculoRepositoryImpl) FindAll() ([]entity.Vehiculo, error) {
	var list []entity.Vehiculo
	if err := database.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *vehiculoRepositoryImpl) FindByID(id uint) (*entity.Vehiculo, error) {
	var v entity.Vehiculo
	if err := database.DB.First(&v, id).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *vehiculoRepositoryImpl) Save(v *entity.Vehiculo) (*entity.Vehiculo, error) {
	if err := database.DB.Create(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

// 👇 el Update que te falta
func (r *vehiculoRepositoryImpl) Update(v *entity.Vehiculo) (*entity.Vehiculo, error) {
	if err := database.DB.Save(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func (r *vehiculoRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&entity.Vehiculo{}, id).Error
}

func (r *vehiculoRepositoryImpl) FindByPlaca(placa string) (*entity.Vehiculo, error) {
	var v entity.Vehiculo
	if err := database.DB.Where("placa = ?", placa).First(&v).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *vehiculoRepositoryImpl) FindByTipo(tipo string) ([]entity.Vehiculo, error) {
	var list []entity.Vehiculo
	if err := database.DB.Where("tipo_vehiculo = ?", tipo).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *vehiculoRepositoryImpl) FindByIdCliente(idCliente uint) ([]entity.Vehiculo, error) {
	var list []entity.Vehiculo
	if err := database.DB.Where("id_cliente = ?", idCliente).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
