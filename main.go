package main

import (
	"log"

	Users "go-template-api/internal/api/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	v1Router := router.Group("/v1")

	Users.RegisterRoutes(v1Router)

	router.Run("localhost:8080")

}
