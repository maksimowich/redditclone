package users_memory_repo

import (
	"github.com/google/uuid"
)

func (repo *UsersMemoryRepo) GetById(
	userId uuid.UUID,
) (*User, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	for _, user := range repo.Users {
		if user.Id == userId {
			return user, nil
		}
	}

	return nil, &UserNotFoundError{
		UserId: userId,
	}
}
