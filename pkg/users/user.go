package users

import "github.com/google/uuid"

type UserAdd struct {
	Username string `json:"username"`
	Password string
}

type User struct {
	Id uuid.UUID `json:"id"`
	UserAdd
	Created string
}

type UsersRepoInterface interface {
	Add(postAdd *UserAdd) (*User, error)
	CheckPassword(username, password string) (*User, error)
}
