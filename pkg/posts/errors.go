package posts

import (
	"fmt"

	"github.com/google/uuid"
)

type PostNotFoundError struct {
	PostID uuid.UUID
}

func (e *PostNotFoundError) Error() string {
	return fmt.Sprintf("post with ID %v not found", e.PostID)
}
