package city

import (
	"gorm.io/gorm"
)

type CityRepository interface {
	FindAll() []City
	FindById(id int) City
	FindByProvinceID(ProvinceID int) []City
}

type CityRepositoryImpl struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &CityRepositoryImpl{db}
}

func (cr *CityRepositoryImpl) FindAll() []City {
	var cities []City

	_ = cr.db.Find(&cities)

	return cities
}

func (cr *CityRepositoryImpl) FindById(id int) City {
	var city City

	_ = cr.db.First(&city, id)

	return city
}

func (cr *CityRepositoryImpl) FindByProvinceID(provinceID int) []City {
	var cities []City

	_ = cr.db.Where("province_id = ?", provinceID).Find(&cities)

	return cities
}
