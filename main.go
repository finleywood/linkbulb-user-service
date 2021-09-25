package main

import (
	"log"

	"github.com/joho/godotenv"
	"linkbulb.io/users-service/models"
	"linkbulb.io/users-service/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.Init()
	routes.Run()
}
