package utils_handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetPostIdFromQuery(
	r *http.Request,
) (uuid.UUID, error) {
	vars := mux.Vars(r)
	postIdStr := vars["postId"]
	postId, err := uuid.Parse(postIdStr)
	return postId, err
}
