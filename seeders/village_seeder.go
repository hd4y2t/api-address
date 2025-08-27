package seeder

import (
	"api-address/moduls/village"
	"fmt"

	"gorm.io/gorm"
)

func VillageSeed(db *gorm.DB) {
	// --- Seed Villages ---
	villages := []village.Village{
		{Code:"001",Name: "LATIUNG", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"002",Name: "LABUHAN BAJAU", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"003",Name: "SUAK LAMATAN", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"004",Name: "ANA AO", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"005",Name: "LATALING", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"006",Name: "PULAU BENGKALAK", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"007",Name: "BADEGONG", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"008",Name: "KEBUN BARU", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"009",Name: "ULUL MAYANG", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"010",Name: "PASIR TINGGI", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"011",Name: "LABUHAN JAYA", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"012",Name: "LABUHAN BAKTI", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"013",Name: "BATU RALANG", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"014",Name: "ALUS ALUS", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"015",Name: "SEUNEUBOK", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"016",Name: "BLANG SEBEL", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"017",Name: "TRANS BARU", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"018",Name: "TRANS MERANTI", ProvinceId: "1", CityId: "1", DistrictId: "1"},
		{Code:"019",Name: "TRANS JERNGE", ProvinceId: "1", CityId: "1", DistrictId: "1"},
	}

	for _, v := range villages {
		db.FirstOrCreate(&v, village.Village{Name: v.Name, ProvinceId: v.ProvinceId, CityId: v.CityId, DistrictId: v.DistrictId})
	}
	fmt.Println("Village seeder berhasil")
}
