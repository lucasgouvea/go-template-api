package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	Guns "go-api/internal/api/guns"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	v1Router := router.Group("/v1")

	Guns.Start(v1Router)

	Setup()

	router.Run("localhost:8080")

}
