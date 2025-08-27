package seeder

import (
	district "api-address/modules/district"
	"fmt"
	"gorm.io/gorm"
)

func DistrictSeed(db *gorm.DB) {
	districts := []district.District{
		{Name: "Jakarta", ProvinceID: 1, CityId: 1},
		{Name: "Bandung", ProvinceID: 1, CityId: 2},
		{Name: "Surabaya", ProvinceID: 1, CityId: 3},
	}

	for _, d := range districts {
		db.FirstOrCreate(&d, district.District{Name: d.Name, ProvinceID: d.ProvinceID, CityId: d.CityId})
	}

	fmt.Println("District seeder berhasil")
}
