package main

import (
	Redis "go-api/internal/redis"
)

func Setup() {

	Redis.NewPool()

}
