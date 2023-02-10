package shared

import (
	"os"
)

// Constants for environment
type GoEnv struct {
	PROD string
	DEV  string
}

// Env var data
type Env struct {
	DB_HOST               string
	DB_USER               string
	DB_PASSWORD           string
	DB_NAME               string
	GO_ENV                string
	DEFAULT_USER_PASSWORD string
}

func getGoEnv() GoEnv {
	return GoEnv{
		PROD: "prod",
		DEV:  "dev",
	}
}

func GetEnvVars() Env {
	return Env{
		DB_HOST:               os.Getenv("DB_HOST"),
		DB_USER:               os.Getenv("DB_USER"),
		DB_PASSWORD:           os.Getenv("DB_PASSWORD"),
		DB_NAME:               os.Getenv("DB_NAME"),
		GO_ENV:                os.Getenv("GO_ENV"),
		DEFAULT_USER_PASSWORD: os.Getenv("DEFAULT_USER_PASSWORD"),
	}
}

func IsProdEnv() bool {
	return GetEnvVars().GO_ENV == getGoEnv().PROD
}
