package users_memory_repo

func (repo *UsersMemoryRepo) CheckPassword(
	username, password string,
) (*User, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	user, ok := repo.Users[username]
	if !ok {
		return nil, &InvalidUsernameOrPasswordtsError{}
	}

	if user.Password != password {
		return nil, &InvalidUsernameOrPasswordtsError{}
	}

	return user, nil
}
