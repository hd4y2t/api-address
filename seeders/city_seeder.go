package seeder

import (
	"api-address/modules/city"
	"fmt"
	"gorm.io/gorm"
)

func SeedCity(db *gorm.DB) {
	cities := []city.City{
		{Name: "Jakarta", ProvinceID: 1},
		{Name: "Bandung", ProvinceID: 2},
		{Name: "Surabaya", ProvinceID: 3},
	}

	for _, c := range cities {
		db.FirstOrCreate(&c, city.City{Name: c.Name, ProvinceID: c.ProvinceID})
	}

	fmt.Println("City seeder berhasil")
}
