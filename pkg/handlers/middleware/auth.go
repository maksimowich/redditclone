package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
	"github.com/maksimowich/redditclone/pkg/jwt"
	"github.com/maksimowich/redditclone/pkg/tokens"
)

func extractToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	if header == "" {
		return ""
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

type contextKeyUserId string
type contextKeyUsername string

const CtxKeyUserId = contextKeyUserId("userId")
const CtxKeyUserName = contextKeyUsername("username")

func AuthMiddleware(
	tokensRepo tokens.TokensRepoInterface,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractToken(r)
			if token == "" {
				err := fmt.Errorf("unauthorized. Access token is not set")
				utils.HandleError(w, http.StatusUnauthorized, err)
				return
			}

			isValid, err := tokensRepo.IsValid(token)
			if err != nil || !isValid {
				err := fmt.Errorf("unauthorized. Access token is invalid")
				utils.HandleError(w, http.StatusUnauthorized, err)
				return
			}

			claims, err := jwt.ParseJWT(token)
			if err != nil {
				err := fmt.Errorf("unauthorized. Failed to parse access token")
				utils.HandleError(w, http.StatusUnauthorized, err)
				return
			}

			ctx := context.WithValue(r.Context(), CtxKeyUserId, claims.UserId)
			ctx = context.WithValue(ctx, CtxKeyUserName, claims.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIdFromContext(
	w http.ResponseWriter,
	r *http.Request,
) (uuid.UUID, error) {
	userIdString, ok := r.Context().Value(CtxKeyUserId).(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user ID not found in context or invalid type")
	}

	userIdUUID, err := uuid.Parse(userIdString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID format in context: %w", err)
	}

	return userIdUUID, nil
}
