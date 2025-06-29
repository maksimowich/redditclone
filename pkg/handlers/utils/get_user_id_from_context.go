package utils_handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/middleware"
)

func GetUserIdFromContext(
	w http.ResponseWriter,
	r *http.Request,
) (uuid.UUID, error) {
	userIdString, ok := r.Context().Value(middleware.CtxKeyUserId).(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user ID not found in context or invalid type")
	}

	userIdUUID, err := uuid.Parse(userIdString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID format in context: %w", err)
	}

	return userIdUUID, nil
}
