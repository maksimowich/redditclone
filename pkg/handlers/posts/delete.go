package posts_handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/handlers/middleware"
	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
)

type DeleteResponseBody struct {
	PostId uuid.UUID `json:"post_id"`
}

func (h *PostsHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetPostIdFromQuery(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	userId, err := middleware.GetUserIdFromContext(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	deletedPostId, err := h.PostsRepo.Delete(postId, userId)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	responseBody := &DeleteResponseBody{
		PostId: deletedPostId,
	}
	utils.JSONResponse(w, http.StatusOK, responseBody)
}
