package main

import (
	Guns "go-api/internal/api/guns"
	Redis "go-api/internal/redis"
)

func Setup() {

	Redis.NewPool()
	Redis.CreateMany(Guns.Guns)

}
