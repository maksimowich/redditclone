package posts_handlers

import "github.com/google/uuid"

type PostIdResponseBody struct {
	PostId uuid.UUID `json:"post_id"`
}

type CommentIdResponseBody struct {
	PostId uuid.UUID `json:"post_id"`
}
