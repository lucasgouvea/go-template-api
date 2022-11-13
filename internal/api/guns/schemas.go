package guns

import "strings"

type IGun struct {
	Id    string  `redis:"id" json:"id"`
	Name  string  `redis:"name" json:"name"`
	Price float64 `redis:"price" json:"price"`
}

func (gun IGun) GetHash() string {
	var hash strings.Builder
	hash.WriteString("gun:")
	hash.WriteString(gun.Id)
	return hash.String()
}
