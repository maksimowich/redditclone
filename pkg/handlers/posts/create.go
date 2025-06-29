package posts_handlers

import (
	"net/http"

	"github.com/maksimowich/redditclone/pkg/handlers/middleware"
	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
)

type AddRequestBody struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	Url      string `json:"url,omitempty"`
}

func (h *PostsHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	addRequestBody := &AddRequestBody{}
	if err := middleware.ValidateRequest(r, addRequestBody); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	userId, err := middleware.GetUserIdFromContext(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	postAdd := &PostAdd{
		Category: addRequestBody.Category,
		Title:    addRequestBody.Title,
		Type:     addRequestBody.Type,
		Text:     addRequestBody.Text,
		Url:      addRequestBody.Url,
		AuthorId: userId,
	}
	createdPost, err := h.PostsRepo.Add(postAdd)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err)
		return
	}

	responseBody := &PostIdResponseBody{createdPost.Id}
	utils.JSONResponse(w, http.StatusCreated, responseBody)
}
