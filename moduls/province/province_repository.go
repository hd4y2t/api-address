package province

import (
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	FindAll() []Province
	FindById(id int) Province
}

type ProvinceRepositoryImpl struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &ProvinceRepositoryImpl{db}
}

func (pr *ProvinceRepositoryImpl) FindAll() []Province {
	var provinces []Province

	_ = pr.db.Find(&provinces)

	return provinces
}

func (pr *ProvinceRepositoryImpl) FindById(id int) Province {
	var province Province

	_ = pr.db.First(&province, id)

	return province
}
