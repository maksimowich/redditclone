package posts_handlers

import (
	"net/http"

	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
)

func (h *PostsHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	var posts []*Post
	var err error
	if category == "" {
		posts, err = h.PostsRepo.GetAll()
		if err != nil {
			utils.HandleError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		posts, err = h.PostsRepo.GetByCategory(category)
		if err != nil {
			utils.HandleError(w, http.StatusInternalServerError, err)
			return
		}
	}

	responseBody := make([]*PostResponse, 0, len(posts))
	for _, post := range posts {
		responseBody = append(responseBody, post.ToResponse())
	}
	utils.JSONResponse(w, http.StatusOK, responseBody)
}

func (h *PostsHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.GetPostIdFromQuery(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	post, err := h.PostsRepo.GetById(postId)
	if err != nil {
		utils.HandleErrorByType(w, err)
		return
	}

	responseBody := post.ToResponse()
	utils.JSONResponse(w, http.StatusOK, responseBody)
}
