package guns

type Gun struct {
	Id    string  `redis:"id"  binding:"required"`
	Name  string  `redis:"name" binding:"required"`
	Price float64 `redis:"price" binding:"required"`
}
