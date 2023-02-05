package gun_schemas

import (
	GunModels "go-template-api/internal/api/guns/models"
)

type GunResponseSchema struct {
	CreatedAt    float64 `json:"created_at"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	SerialNumber string  `json:"serial_number"`
}

func NewGunResponseSchema(model GunModels.GunModel) GunResponseSchema {
	schema := new(GunResponseSchema)
	schema.CreatedAt = model.Data.CreatedAt
	schema.Name = model.Data.Name
	schema.Price = model.Data.Price
	schema.SerialNumber = model.Data.SerialNumber
	return *schema
}
