package users

import (
	Database "go-template-api/internal/database"
	Errors "go-template-api/internal/errors"

	"gorm.io/gorm/clause"
)

func listUsers() ([]User, error) {
	users := []User{}
	db := Database.GetDB()
	err := db.Select("id", "created_at", "name").Find(&users).Error
	return users, err
}

func createUser(user *User) error {
	user.hashPassword()
	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&user).Error
	return err
}

func updateUser(user *User) error {
	user.hashPassword()
	db := Database.GetDB()
	res := db.Clauses(clause.Returning{}).Where("id = ?", user.ID).Updates(user)
	if res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
