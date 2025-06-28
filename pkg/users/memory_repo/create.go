package users_memory_repo

import (
	"time"

	"github.com/google/uuid"
)

func (repo *UsersMemoryRepo) Add(
	userAdd *UserAdd,
) (*User, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	username := userAdd.Username

	_, ok := repo.Users[username]
	if ok {
		return nil, &UserAlreadyExistsError{
			Username: username,
		}
	}

	newUser := &User{
		Id:       uuid.New(),
		Username: username,
		Password: userAdd.Password,
		Created:  time.Now().String(),
	}
	repo.Users[username] = newUser

	return newUser, nil
}
