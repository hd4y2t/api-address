package cmd

import (
	"api-address/database"
	"api-address/seeders"
	"fmt"
)

func SeedMain() {
	db := database.DB()

	seeder.SeedAll(db)

	fmt.Println("🎉 Semua seeder berhasil dijalankan!")
}
