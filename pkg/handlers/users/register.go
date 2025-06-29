package users_handlers

import (
	"errors"
	"net/http"

	"github.com/maksimowich/redditclone/pkg/handlers/middleware"
	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
	"github.com/maksimowich/redditclone/pkg/jwt"
	"github.com/maksimowich/redditclone/pkg/users"
)

type RegisterRequestBody struct {
	Username string `json:"username" schema:"username" valid:"required"`
	Password string `json:"password" schema:"password" valid:"required"`
}

type RegisterResponseBody struct {
	Token string `json:"token" schema:"token" valid:"required"`
}

func (h *UsersHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	registerRequestBody := &RegisterRequestBody{}

	if err := middleware.ValidateRequest(r, registerRequestBody); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	userAdd := &users.UserAdd{
		Username: registerRequestBody.Username,
		Password: registerRequestBody.Password,
	}

	user, err := h.UsersRepo.Add(userAdd)
	var UserAlreadyExistsErr *users.UserAlreadyExistsError
	if err != nil && errors.As(err, &UserAlreadyExistsErr) {
		utils.HandleError(w, http.StatusConflict, err)
		return
	} else if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err)
		return
	}

	token, expires, err := jwt.GenerateJWT(user.Id, user.Username)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err)
		return
	}
	err = h.TokensRepo.Add(token, expires)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err)
		return
	}

	registerResponseBody := &RegisterResponseBody{
		Token: token,
	}
	utils.JSONResponse(w, http.StatusOK, registerResponseBody)
}
