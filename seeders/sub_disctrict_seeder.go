package seeder

import (
	subDistrict "api-address/modules/sub_district"
	"fmt"
	"gorm.io/gorm"
)

func SubDistrictSeed(db *gorm.DB) {
	subDistricts := []subDistrict.SubDistrict{
		{Name: "Jakarta", ProvinceID: 1, CityId: 1},
		{Name: "Bandung", ProvinceID: 1, CityId: 2},
		{Name: "Surabaya", ProvinceID: 1, CityId: 3},
	}

	for _, sd := range subDistricts {
		db.FirstOrCreate(&sd, subDistrict.SubDistrict{Name: sd.Name, ProvinceID: sd.ProvinceID, CityId: sd.CityId})
	}

	fmt.Println("Sub District seeder berhasil")
}
