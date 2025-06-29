package posts_handlers

import (
	"net/http"

	"github.com/maksimowich/redditclone/pkg/handlers/middleware"
	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
)

type AddCommentRequestBody struct {
	Text string `json:"text"`
}

func (h *PostsHandler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetPostIdFromQuery(r)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	addCommentRequestBody := &AddCommentRequestBody{}
	if err := middleware.ValidateRequest(r, addCommentRequestBody); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	userId, err := middleware.GetUserIdFromContext(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	newCommentId, err := h.PostsRepo.AddComment(postId, addCommentRequestBody.Text, userId)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	responseBody := &CommentIdResponseBody{newCommentId}
	utils.JSONResponse(w, http.StatusOK, responseBody)
}

func (h *PostsHandler) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetPostIdFromQuery(r)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	commentId, err := utils.GetCommentIdFromQuery(r)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	userId, err := middleware.GetUserIdFromContext(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	deletedCommentId, err := h.PostsRepo.DeleteComment(postId, commentId, userId)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	responseBody := &CommentIdResponseBody{deletedCommentId}
	utils.JSONResponse(w, http.StatusOK, responseBody)
}

func (h *PostsHandler) UpvoteHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetPostIdFromQuery(r)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	userId, err := middleware.GetUserIdFromContext(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	upvotedPostId, err := h.PostsRepo.Upvote(postId, userId)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	responseBody := &PostIdResponseBody{upvotedPostId}
	utils.JSONResponse(w, http.StatusOK, responseBody)
}

func (h *PostsHandler) DownvoteHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetPostIdFromQuery(r)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	userId, err := middleware.GetUserIdFromContext(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	upvotedPostId, err := h.PostsRepo.Downvote(postId, userId)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	responseBody := &PostIdResponseBody{upvotedPostId}
	utils.JSONResponse(w, http.StatusOK, responseBody)
}
