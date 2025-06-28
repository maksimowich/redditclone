package users

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string
	Created  string
}

type UserAdd struct {
	Username string `json:"username"`
	Password string
}

type UsersRepoInterface interface {
	// Create
	Add(userAdd *UserAdd) (*User, error)

	// Read
	GetById(userId uuid.UUID) (*User, error)

	// Others
	CheckPassword(username, password string) (*User, error)
}
