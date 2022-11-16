package gun_schemas

import GunModels "go-api/internal/api/guns/models"

type GunResponseSchema struct {
	CreatedAt    float64 `json:"created_at"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	SerialNumber string  `json:"serial_number"`
}

func NewGunResponseSchema(model *GunModels.GunModel) GunResponseSchema {
	gunSchema := new(GunResponseSchema)
	gunSchema.CreatedAt = model.Data.CreatedAt
	gunSchema.Name = model.Data.Name
	gunSchema.Price = model.Data.Price
	gunSchema.SerialNumber = model.Data.SerialNumber
	return *gunSchema
}
