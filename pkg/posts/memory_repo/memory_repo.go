package posts_memory_repo

import (
	"sync"

	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/posts"
	"github.com/maksimowich/redditclone/pkg/users"
)

type (
	Post              = posts.Post
	PostAdd           = posts.PostAdd
	PostNotFoundError = posts.PostNotFoundError
	Vote              = posts.Vote
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
