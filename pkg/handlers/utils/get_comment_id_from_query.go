package utils_handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetCommentIdFromQuery(
	r *http.Request,
) (uuid.UUID, error) {
	vars := mux.Vars(r)
	commentIdStr := vars["commentId"]
	commentId, err := uuid.Parse(commentIdStr)
	return commentId, err
}
