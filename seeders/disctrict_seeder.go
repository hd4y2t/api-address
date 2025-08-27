package seeder

import (
	district "api-address/moduls/district"
	"fmt"
	"gorm.io/gorm"
)

func DistrictSeed(db *gorm.DB) {
	districts := []district.District{
		{Code: "011", Name: "TEUPAH SELATAN", ProvinceId: "1", CityId: "1"},
		{Code: "020", Name: "SIMEULUE TIMUR", ProvinceId: "1", CityId: "1"},
		{Code: "021", Name: "TEUPAH BARAT", ProvinceId: "1", CityId: "1"},
		{Code: "022", Name: "TEUPAH TENGAH", ProvinceId: "1", CityId: "1"},
		{Code: "030", Name: "SIMEULUE TENGAH", ProvinceId: "1", CityId: "1"},
		{Code: "031", Name: "TELUK DALAM", ProvinceId: "1", CityId: "1"},
		{Code: "032", Name: "SIMEULUE CUT", ProvinceId: "1", CityId: "1"},
		{Code: "040", Name: "SALANG", ProvinceId: "1", CityId: "1"},
		{Code: "050", Name: "SIMEULUE BARAT", ProvinceId: "1", CityId: "1"},
		{Code: "051", Name: "ALAFAN", ProvinceId: "1", CityId: "1"},
	}

	for _, d := range districts {
		db.FirstOrCreate(&d, district.District{Name: d.Name, ProvinceId: d.ProvinceId, CityId: d.CityId})
	}

	fmt.Println("District seeder berhasil")
}
