package users

import (
	Shared "go-template-api/internal/shared"
)

const DEFAULT_USER_NAME = "admin"

func Seed() {
	user := User{Name: DEFAULT_USER_NAME, Password: Shared.GetEnvVars().DEFAULT_USER_PASSWORD}
	createUser(&user)
}
