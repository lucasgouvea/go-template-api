package gun_schemas

import GunModels "go-template-api/internal/api/guns/models"

type GunPostSchema struct {
	Name         string  `json:"name" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	SerialNumber string  `json:"serial_number" binding:"required"`
}

func (schema GunPostSchema) GetGun() GunModels.Gun {
	var gun = GunModels.Gun{Name: schema.Name, Price: schema.Price, SerialNumber: schema.SerialNumber}
	return gun
}
