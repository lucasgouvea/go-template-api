package main

import (
	"fmt"
	"os"

	Users "go-template-api/internal/api/users"
	Database "go-template-api/internal/database"
	Shared "go-template-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	fmt.Printf("Execution args: %v\n", args)

	if !Shared.IsProdEnv() {
		fmt.Println("Loading .env")
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	Database.Start(Shared.GetEnvVars().DB_HOST, Shared.GetEnvVars().DB_USER, Shared.GetEnvVars().DB_PASSWORD, Shared.GetEnvVars().DB_NAME)
	if err := migrate(); err != nil {
		panic(err)
	}

	if len(args) > 1 {
		if args[1] == "seed" {
			seed()
		}
	} else {
		if err := startAPI(); err != nil {
			panic(err)
		}
	}
}

func startAPI() error {
	router := gin.Default()

	v1Router := router.Group("/v1")

	Users.RegisterRoutes(v1Router)
	Drivers.RegisterRoutes(v1Router)

	return router.Run("0.0.0.0:8081")
}

func migrate() error {
	fmt.Println(" *** Running migrations ***")
	models := []any{&Users.User{}, &Drivers.Driver{}}
	return Database.Migrate(models)
}

func seed() {
	Users.Seed()
}
