package main

import (
	"fmt"
	"household-dashboard/src/config"
	"household-dashboard/src/database"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	fmt.Println("Hello, World!")
}
