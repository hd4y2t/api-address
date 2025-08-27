package sub_district

import (
	"gorm.io/gorm"
)

type SubDistrictRepository interface {
	FindAll() []SubDistrict
	FindById(id int) SubDistrict
	FindByProvinceID(provinceID int) []SubDistrict
	FindByCityID(cityID int) []SubDistrict
}

type SubDistrictRepositoryImpl struct {
	db *gorm.DB
}

func NewSubDistrictRepository(db *gorm.DB) SubDistrictRepository {
	return &SubDistrictRepositoryImpl{db}
}

func (pr *SubDistrictRepositoryImpl) FindAll() []SubDistrict {
	var subDistricts []SubDistrict

	_ = pr.db.Find(&subDistricts)

	return subDistricts
}

func (pr *SubDistrictRepositoryImpl) FindById(id int) SubDistrict {
	var subDistrict SubDistrict

	_ = pr.db.First(&subDistrict, id)

	return subDistrict
}

func (pr *SubDistrictRepositoryImpl) FindByProvinceID(provinceID int) []SubDistrict {
	var subDistricts []SubDistrict

	_ = pr.db.Where("province_id = ?", provinceID).Find(&subDistricts)

	return subDistricts
}

func (pr *SubDistrictRepositoryImpl) FindByCityID(cityID int) []SubDistrict {
	var subDistricts []SubDistrict

	_ = pr.db.Where("city_id = ?", cityID).Find(&subDistricts)

	return subDistricts
}
