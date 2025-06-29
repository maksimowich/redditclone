package tokens

import (
	"time"
)

type TokensRepoInterface interface {
	// Create
	Add(token string, expires time.Time) error

	// Delete
	Delete(token string) error

	// Others
	IsValid(token string) (bool, error)
}
