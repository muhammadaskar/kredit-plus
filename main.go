package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadaskar/kredit-plus/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	router := routes.SetupRouter()
	router.Run(":" + port)
}
