package main

import (
	"log"

	Guns "go-template-api/internal/api/guns"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Setup()

	router := gin.Default()

	v1Router := router.Group("/v1")

	Guns.Start(v1Router)

	router.Run("localhost:8080")

}
