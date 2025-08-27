package district

import (
	"gorm.io/gorm"
)

type DistrictRepository interface {
	FindAll() []District
	FindById(id int) District
	FindByProvinceID(provinceID int) []District
	FindByCityID(cityID int) []District
}

type DistrictRepositoryImpl struct {
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &DistrictRepositoryImpl{db}
}

func (pr *DistrictRepositoryImpl) FindAll() []District {
	var districts []District

	_ = pr.db.Find(&districts)

	return districts
}

func (pr *DistrictRepositoryImpl) FindById(id int) District {
	var district District

	_ = pr.db.First(&district, id)

	return district
}

func (pr *DistrictRepositoryImpl) FindByProvinceID(provinceID int) []District {
	var districts []District

	_ = pr.db.Where("province_id = ?", provinceID).Find(&districts)

	return districts
}

func (pr *DistrictRepositoryImpl) FindByCityID(cityID int) []District {
	var districts []District

	_ = pr.db.Where("city_id = ?", cityID).Find(&districts)

	return districts
}
