package posts_memory_repo

import (
	"sync"

	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/posts"
	"github.com/maksimowich/redditclone/pkg/users"
)

type (
	Vote    = posts.Vote
	Comment = posts.Comment
	Post    = posts.Post

	PostAdd = posts.PostAdd

	CommentNotFoundError   = posts.CommentNotFoundError
	PostNotFoundError      = posts.PostNotFoundError
	UserHasNotEnoughRights = posts.UserHasNotEnoughRights
)

type PostsMemoryRepo struct {
	Posts     map[uuid.UUID]*Post
	Mu        *sync.Mutex
	UsersRepo users.UsersRepoInterface
}

func NewPostsMemoryRepo(
	usersRepo users.UsersRepoInterface,
) *PostsMemoryRepo {
	return &PostsMemoryRepo{
		Posts:     make(map[uuid.UUID]*Post),
		Mu:        &sync.Mutex{},
		UsersRepo: usersRepo,
	}
}
