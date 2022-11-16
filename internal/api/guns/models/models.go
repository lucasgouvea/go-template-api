package gun_models

import (
	Redis "go-api/internal/redis"
	"strings"
	"time"
)

type Gun struct {
	CreatedAt    float64 `redis:"created_at"`
	Name         string  `redis:"name"`
	Price        float64 `redis:"price"`
	SerialNumber string  `redis:"serial_number"`
}

type GunModel = Redis.Model[Gun]

const GunModelName = "guns"

func NewGunModel(gun Gun) *GunModel {
	gunModel := new(GunModel)
	gunModel.Data.CreatedAt = float64(time.Now().Unix())
	gunModel.Data.Name = gun.Name
	gunModel.Data.Price = gun.Price
	gunModel.Data.SerialNumber = gun.SerialNumber
	gunModel.Meta.DefaultScore = gunModel.Data.CreatedAt
	gunModel.Data.Name = GunModelName
	gunModel.Meta.Hash = strings.Join([]string{GunModelName, ":", gun.SerialNumber}, "")
	gunModel.Meta.SortedSet = GunModelName
	return gunModel
}
