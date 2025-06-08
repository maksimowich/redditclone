package users

import (
	"fmt"
)

type UserAlreadyExistsError struct {
	Username string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with username %v already exists", e.Username)
}

type InvalidUsernameOrPasswordtsError struct{}

func (e *InvalidUsernameOrPasswordtsError) Error() string {
	return "invalid username or password"
}
