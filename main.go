package main

import (
	Guns "go-api/internal/api/guns"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1Router := router.Group("/v1")

	Guns.Start(v1Router)

	router.Run("localhost:8080")
}
