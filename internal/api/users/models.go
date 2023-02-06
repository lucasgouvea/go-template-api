package users

import "time"

type User struct {
	ID        uint
	CreatedAt time.Time
	Name      string
}
