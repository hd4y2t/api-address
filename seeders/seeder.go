package seeder

import "gorm.io/gorm"

func SeedAll(db *gorm.DB) {
	ProvinceSeed(db)
	CitySeed(db)
	DistrictSeed(db)
	VillageSeed(db)
}
