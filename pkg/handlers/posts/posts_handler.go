package posts_handlers

import (
	"github.com/maksimowich/redditclone/pkg/posts"
)

type (
	Post         = posts.Post
	PostAdd      = posts.PostAdd
	PostResponse = posts.PostResponse
)

type PostsHandler struct {
	PostsRepo posts.PostsRepoInterface
}
