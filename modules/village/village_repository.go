package village

import (
	"gorm.io/gorm"
)

type VillageRepository interface {
	FindAll() []Village
	FindById(id int) Village
	FindByProvinceID(ProvinceID int) []Village
	FindByCityID(CityID int) []Village
	FindBySubDistrictID(SubDistrictID int) []Village
}

type VillageRepositoryImpl struct {
	db *gorm.DB
}

func NewVillageRepository(db *gorm.DB) VillageRepository {
	return &VillageRepositoryImpl{db}
}

func (vr *VillageRepositoryImpl) FindAll() []Village {
	var villages []Village

	_ = vr.db.Find(&villages)

	return villages
}

func (vr *VillageRepositoryImpl) FindById(id int) Village {
	var village Village

	_ = vr.db.First(&village, id)

	return village
}

func (vr *VillageRepositoryImpl) FindByProvinceID(provinceID int) []Village {
	var villages []Village

	_ = vr.db.Where("province_id = ?", provinceID).Find(&villages)

	return villages
}

func (vr *VillageRepositoryImpl) FindByCityID(cityID int) []Village {
	var villages []Village

	_ = vr.db.Where("city_id = ?", cityID).Find(&villages)

	return villages
}

func (vr *VillageRepositoryImpl) FindBySubDistrictID(subDistrictID int) []Village {
	var villages []Village

	_ = vr.db.Where("sub_district_id = ?", subDistrictID).Find(&villages)

	return villages
}
