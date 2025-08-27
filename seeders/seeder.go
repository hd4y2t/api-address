package seeder

import (
	"api-address/modules/city"
	"api-address/modules/province"
	subDistrict "api-address/modules/sub_discrict"
	"api-address/modules/village"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// --- Seed Provinces ---
	provinces := []province.Province{
		{Name: "DKI Jakarta"},
		{Name: "Jawa Barat"},
		{Name: "Jawa Timur"},
	}

	for _, p := range provinces {
		db.FirstOrCreate(&p, province.Province{Name: p.Name})
	}

	cities := []city.City{
		{Name: "Jakarta", ProvinceID: 1},
		{Name: "Bandung", ProvinceID: 2},
		{Name: "Surabaya", ProvinceID: 3},
	}

	for _, c := range cities {
		db.FirstOrCreate(&c, city.City{Name: c.Name, ProvinceID: c.ProvinceID})
	}

	

	// --- Seed Villages ---
	villages := []village.Village{
		{Name: "Kebayoran", ProvinceID: 1, CityId: 1, SubDistrictID: 1},
		{Name: "Menteng", ProvinceID: 1, CityId: 1, SubDistrictID: 1},
		{Name: "Dago", ProvinceID: 1, CityId: 2, SubDistrictID: 2},
		{Name: "Wonokromo", ProvinceID: 1, CityId: 3, SubDistrictID: 3},
	}

	for _, v := range villages {
		db.FirstOrCreate(&v, village.Village{Name: v.Name, ProvinceID: v.ProvinceID, CityId: v.CityId, SubDistrictID: v.SubDistrictID})
	}
}
