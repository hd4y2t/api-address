package main

import (
    "api-address/database" // harus sama persis dengan nama module di go.mod
    "api-address/seeders"
    "fmt"
)

func main() {
    db := database.DB()
    seeder.SeedAll(db)
    fmt.Println("🎉 Semua seeder berhasil dijalankan!")
}