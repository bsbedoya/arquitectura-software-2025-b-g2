package repository

import (
	"crud-all-project/database"
	"crud-all-project/entity"

	"gorm.io/gorm"
)

type servicioRepositoryImpl struct {
	db *gorm.DB
}

func NewServicioRepository() ServicioRepository {
	return &servicioRepositoryImpl{db: database.DB}
}

func (r *servicioRepositoryImpl) FindAll() ([]entity.Servicio, error) {
	var items []entity.Servicio
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *servicioRepositoryImpl) FindByID(id uint) (*entity.Servicio, error) {
	var s entity.Servicio
	if err := r.db.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *servicioRepositoryImpl) Save(s *entity.Servicio) (*entity.Servicio, error) {
	if err := r.db.Save(s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func (r *servicioRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.Servicio{}, id).Error
}

func (r *servicioRepositoryImpl) FindByNombre(nombre string) (*entity.Servicio, error) {
	var s entity.Servicio
	if err := r.db.Where("nombre_servicio = ?", nombre).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
