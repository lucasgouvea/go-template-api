package guns

import (
	Redis "go-api/internal/redis"
	"reflect"
	"strings"
)

var ModelName = "gun"

type Gun struct {
	Name       string  `redis:"name" json:"name" binding:"required"`
	Price      float64 `redis:"price" json:"price" binding:"required"`
	SerialCode string  `redis:"serial_code" json:"serial_code" binding:"required"`
}

func (gun Gun) GetModelName() string {
	return reflect.TypeOf(gun).String()
}

func NewGunModel(gun Gun) *Redis.Model[Gun] {
	gunModel := new(Redis.Model[Gun])
	gunModel.Data = gun
	gunModel.DefaultScore = gun.SerialCode
	gunModel.Hash = strings.Join([]string{gun.GetModelName(), ":", gun.SerialCode}, "")
	gunModel.SortedSet = gun.GetModelName()
	return gunModel
}
