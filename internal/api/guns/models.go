package guns

type IGun struct {
	Id    string  `redis:"id" json:"id"`
	Name  string  `redis:"name" json:"name"`
	Price float64 `redis:"price" json:"price"`
}
