package users_handlers

import (
	"github.com/maksimowich/redditclone/pkg/tokens"
	"github.com/maksimowich/redditclone/pkg/users"
)

type UsersHandler struct {
	UsersRepo  users.UsersRepoInterface
	TokensRepo tokens.TokensRepoInterface
}
