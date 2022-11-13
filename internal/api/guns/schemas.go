package guns

type IGun struct {
	Id    string  `redis:"id"`
	Name  string  `redis:"name"`
	Price float64 `redis:"price"`
}
