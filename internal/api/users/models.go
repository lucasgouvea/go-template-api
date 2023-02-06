package users

import "time"

type User struct {
	ID        uint      `json:"id" `
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name" binding:"required"`
}
