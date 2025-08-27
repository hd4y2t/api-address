package seeder

import (
	"api-address/modules/city"
	"fmt"
	"gorm.io/gorm"
)

func CitySeed(db *gorm.DB) {
	cities := []city.City{
		{
			Name: "KAB. SIMEULUE",
			ProvinceID: 1,
		},
		{
			Name: "KAB. ACEH SINGKIL",
			ProvinceID: 1,
		},
		{
			Name: "KAB. ACEH SELATAN",
			ProvinceID: 1,
		},
		{
			Name: "KAB. ACEH TENGGARA",
			ProvinceID: 1,
		},
		{
			Name: "KAB. ACEH TIMUR",
			ProvinceID: 1,
    "06": "KAB. ACEH TENGAH",
    "07": "KAB. ACEH BARAT",
    "08": "KAB. ACEH BESAR",
    "09": "KAB. PIDIE",
    "10": "KAB. BIREUEN",
    "11": "KAB. ACEH UTARA",
    "12": "KAB. ACEH BARAT DAYA",
    "13": "KAB. GAYO LUES",
    "14": "KAB. ACEH TAMIANG",
    "15": "KAB. NAGAN RAYA",
    "16": "KAB. ACEH JAYA",
    "17": "KAB. BENER MERIAH",
    "18": "KAB. PIDIE JAYA",
    "71": "KOTA BANDA ACEH",
    "72": "KOTA SABANG",
    "73": "KOTA LANGSA",
    "74": "KOTA LHOKSEUMAWE",
    "75": "KOTA SUBULUSSALAM",
	}

	for _, c := range cities {
		db.FirstOrCreate(&c, city.City{Name: c.Name, ProvinceID: c.ProvinceID})
	}

	fmt.Println("City seeder berhasil")
}
