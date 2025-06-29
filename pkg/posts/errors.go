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

type UserHasNotEnoughRights struct {
	UserId uuid.UUID
}

func (e *UserHasNotEnoughRights) Error() string {
	return fmt.Sprintf("user with ID %v has not enough rights", e.UserId)
}

type CommentNotFoundError struct {
	CommentId uuid.UUID
}

func (e *CommentNotFoundError) Error() string {
	return fmt.Sprintf("comment with ID %v not found", e.CommentId)
}
