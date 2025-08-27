package seeder

import (
	"api-address/modules/province"
	"fmt"
	"gorm.io/gorm"
)

func ProvinceSeed(db *gorm.DB) {
	provinces := []province.Province{
		{Name: "DKI Jakarta"},
		{Name: "Jawa Barat"},
		{Name: "Jawa Timur"},
	}

	for _, p := range provinces {
		db.FirstOrCreate(&p, province.Province{Name: p.Name})
	}

	fmt.Println("✅ Province seeder berhasil")
}
