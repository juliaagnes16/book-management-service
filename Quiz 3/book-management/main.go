package main

import (
	"log"

	"main.go/database"
	"main.go/routes"
)

func main() {
	database.InitDB()
	router := routes.InitRoutes()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
