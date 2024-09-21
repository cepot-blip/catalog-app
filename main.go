package main

import (
	"log"

	"github.com/cepot-blip/catalog-app/config"
	"github.com/cepot-blip/catalog-app/models"
	"github.com/cepot-blip/catalog-app/routes"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    config.InitDB()

    config.DB.AutoMigrate(&models.Product{})

    r := routes.SetupRouter()

    r.Run(":8080")
}