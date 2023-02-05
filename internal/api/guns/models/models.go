package gun_models

import (
	Redis "go-template-api/internal/redis"
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

func GetGunHash(serialNumber string) string {
	return strings.Join([]string{GunModelName, ":", serialNumber}, "")
}

func NewGunModel(gun Gun) *GunModel {
	model := new(GunModel)
	model.Data.CreatedAt = float64(time.Now().Unix())
	model.Data.Name = gun.Name
	model.Data.Price = gun.Price
	model.Data.SerialNumber = gun.SerialNumber
	model.Meta.DefaultScore = model.Data.CreatedAt
	model.Meta.Hash = GetGunHash(gun.SerialNumber)
	model.Meta.SortedSet = GunModelName
	return model
}
