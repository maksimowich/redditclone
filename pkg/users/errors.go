package users

import (
	"fmt"

	"github.com/google/uuid"
)

type UserAlreadyExistsError struct {
	Username string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with username %v already exists", e.Username)
}

type UserNotFoundError struct {
	UserId uuid.UUID
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with ID %v not found", e.UserId)
}

type InvalidUsernameOrPasswordtsError struct{}

func (e *InvalidUsernameOrPasswordtsError) Error() string {
	return "invalid username or password"
}
