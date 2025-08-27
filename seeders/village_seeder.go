package seeder

import (
	"api-address/modules/village"
	"fmt"

	"gorm.io/gorm"
)

func VillageSeed(db *gorm.DB) {
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
	fmt.Println("Village seeder berhasil")
}
