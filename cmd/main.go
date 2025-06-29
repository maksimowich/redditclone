package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/maksimowich/redditclone/pkg/handlers/middleware"
	posts_handlers "github.com/maksimowich/redditclone/pkg/handlers/posts"
	users_handlers "github.com/maksimowich/redditclone/pkg/handlers/users"
	posts_memory_repo "github.com/maksimowich/redditclone/pkg/posts/memory_repo"
	tokens_memory_repo "github.com/maksimowich/redditclone/pkg/tokens/memory_repo"
	users_memory_repo "github.com/maksimowich/redditclone/pkg/users/memory_repo"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync() // flushes buffer, if any
	logger := zapLogger.Sugar()

	usersRepo := users_memory_repo.NewUsersMemoryRepo()
	tokensRepo := tokens_memory_repo.NewTokensMemoryRepo()
	postsRepo := posts_memory_repo.NewPostsMemoryRepo(usersRepo)

	usersHandler := &users_handlers.UsersHandler{
		UsersRepo:  usersRepo,
		TokensRepo: tokensRepo,
	}
	postsHandler := &posts_handlers.PostsHandler{
		PostsRepo: postsRepo,
	}

	authMiddleware := middleware.AuthMiddleware(tokensRepo)

	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/register", usersHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", usersHandler.LoginHandler).Methods("POST")

	// Posts routes
	r.Handle("/posts", authMiddleware(http.HandlerFunc(postsHandler.AddHandler))).Methods("POST")
	r.Handle("/posts/{postId}", authMiddleware(http.HandlerFunc(postsHandler.DeleteHandler))).Methods("DELETE")
	r.HandleFunc("/posts", postsHandler.GetAllHandler).Methods("GET")
	r.HandleFunc("/posts/{postId}", postsHandler.GetByIdHandler).Methods("GET")
	r.Handle("/posts/{postId}/upvote", authMiddleware(http.HandlerFunc(postsHandler.UpvoteHandler))).Methods("POST")
	r.Handle("/posts/{postId}/downvote", authMiddleware(http.HandlerFunc(postsHandler.DownvoteHandler))).Methods("POST")

	// Posts comments routes
	r.Handle("/posts/{postId}/comments", authMiddleware(http.HandlerFunc(postsHandler.AddCommentHandler))).Methods("POST")
	r.Handle("/posts/{postId}/comments/{commentId}", authMiddleware(http.HandlerFunc(postsHandler.DeleteCommentHandler))).Methods("DELETE")

	mux := middleware.AccessLog(logger, r)

	addr := ":8081"
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	http.ListenAndServe(addr, mux)
}
