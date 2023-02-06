package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "users"
}

func (u *User) hashPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		panic(err)
	}
	u.Password = string(bytes)
}

func (u *User) checkPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	return err == nil
}
