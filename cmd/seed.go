package cmd

import (
	"../seeder"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	seeder.ProvinceSeed(db)
	seeder.CitySeed(db)
	seeder.VillageCity(db)
	seeder.SubDistrictSeed(db)
}
