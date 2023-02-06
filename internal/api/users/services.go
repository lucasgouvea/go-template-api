package users

import (
	Database "go-template-api/internal/database"
)

func CreateUser(user User) error {
	user.hashPassword()
	db := Database.GetDB()
	err := db.Create(&user).Error
	return err
}
