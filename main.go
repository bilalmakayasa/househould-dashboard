package main

import (
	"fmt"
	"household-dashboard/src/config"
	"household-dashboard/src/controller"
	"household-dashboard/src/database"
	"household-dashboard/src/delivery/http"
	"household-dashboard/src/repository"
	service "household-dashboard/src/services"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	initRepository := repository.InitRepositories(database.DB)
	initeService := service.InitServices(initRepository)
	initContrroller := controller.InitControllers(initeService)

	httpInit := http.NewUserHttpHandler(initContrroller)
	server := httpInit.RegisterHttpHandler()

	fmt.Println("Hello, World!")

	server.Run(":8080")
}
