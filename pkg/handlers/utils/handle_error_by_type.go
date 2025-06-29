package utils_handlers

import (
	"errors"
	"net/http"

	"github.com/maksimowich/redditclone/pkg/posts"
)

func HandleErrorByType(
	w http.ResponseWriter,
	err error,
) {
	var postNotFoundError *posts.PostNotFoundError
	var userHasNotEnoughRights *posts.UserHasNotEnoughRights
	switch {
	case errors.As(err, &postNotFoundError):
		HandleError(w, http.StatusNotFound, err)
	case errors.As(err, &userHasNotEnoughRights):
		HandleError(w, http.StatusForbidden, err)
	default:
		HandleError(w, http.StatusInternalServerError, err)
	}
}
