package cmd

import (
    "api-address/database"
    "api-address/seeder"
    "fmt"
)

func main() {
    db := database.DB()
    seeder.Seed(db)
    fmt.Println("✅ Seeder berhasil dijalankan!")
}