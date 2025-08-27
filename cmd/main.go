package cmd

import (
	"api-address/database"
	"fmt"
)

func main() {
	db := database.DB()
	fmt.Println("Database connected:", db != nil)
}
