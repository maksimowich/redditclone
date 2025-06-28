package users_memory_repo

import (
	"sync"

	"github.com/maksimowich/redditclone/pkg/users"
)

type (
	User    = users.User
	UserAdd = users.UserAdd

	UserAlreadyExistsError           = users.UserAlreadyExistsError
	UserNotFoundError                = users.UserNotFoundError
	InvalidUsernameOrPasswordtsError = users.InvalidUsernameOrPasswordtsError
)

type UsersMemoryRepo struct {
	Users map[string]*User // username -> *User
	Mu    *sync.Mutex
}

func NewUsersMemoryRepo() *UsersMemoryRepo {
	return &UsersMemoryRepo{
		Users: make(map[string]*User),
		Mu:    &sync.Mutex{},
	}
}
