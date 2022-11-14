package guns

type IGun struct {
	Id    string  `redis:"id" json:"id" binding:"required"`
	Name  string  `redis:"name" json:"name" binding:"required"`
	Price float64 `redis:"price" json:"price" binding:"required"`
}
