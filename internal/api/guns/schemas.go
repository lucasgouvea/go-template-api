package guns

type GunSchema struct {
	Id    string  `json:"id"  binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
