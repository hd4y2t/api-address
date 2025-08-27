package seeder

import (
	"api-address/moduls/city"
	"fmt"
	"gorm.io/gorm"
)

func CitySeed(db *gorm.DB) {
	cities := []city.City{
		{
			Name: "KAB. SIMEULUE",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH SINGKIL",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH SELATAN",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH TENGGARA",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH TIMUR",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH TENGAH",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH BARAT",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH BESAR",
			ProvinceId: "1",
		},
		{
			Name: "KAB. PIDIE",
			ProvinceId: "1",
		},
		{
			Name: "KAB. BIREUEN",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH UTARA",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH BARAT DAYA",
			ProvinceId: "1",
		},
		{
			Name: "KAB. GAYO LUES",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH TAMIANG",
			ProvinceId: "1",
		},
		{
			Name: "KAB. NAGAN RAYA",
			ProvinceId: "1",
		},
		{
			Name: "KAB. ACEH JAYA",
			ProvinceId: "1",
		},
		{
			Name: "KAB. BENER MERIAH",
			ProvinceId: "1",
		},
		{
			Name: "KAB. PIDIE JAYA",
			ProvinceId: "1",
		},
		{
			Name: "KOTA BANDA ACEH",
			ProvinceId: "1",
		},
		{
			Name: "KOTA SABANG",
			ProvinceId: "1",
		},
		{
			Name: "KOTA LANGSA",
			ProvinceId: "1",
		},
		{
			Name: "KOTA LHOKSEUMAWE",
			ProvinceId: "1",
		},
		{
			Name: "KOTA SUBULUSSALAM",
			ProvinceId: "1",
		},
	}

	for _, c := range cities {
		db.FirstOrCreate(&c, city.City{Name: c.Name, ProvinceId: c.ProvinceId})
	}

	fmt.Println("City seeder berhasil")
}
