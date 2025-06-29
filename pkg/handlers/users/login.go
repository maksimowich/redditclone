package users_handlers

import (
	"net/http"

	"github.com/maksimowich/redditclone/pkg/handlers/middleware"
	utils "github.com/maksimowich/redditclone/pkg/handlers/utils"
	"github.com/maksimowich/redditclone/pkg/jwt"
)

type LoginRequestBody struct {
	Username string `json:"username" schema:"username" valid:"required"`
	Password string `json:"password" schema:"password" valid:"required"`
}

type LoginResponseBody struct {
	Token string `json:"token" schema:"token" valid:"required"`
}

func (h *UsersHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginRequestBody := &LoginRequestBody{}

	if err := middleware.ValidateRequest(r, loginRequestBody); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	username := loginRequestBody.Username
	password := loginRequestBody.Password

	user, err := h.UsersRepo.CheckPassword(username, password)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err)
		return
	}

	token, expires, err := jwt.GenerateJWT(user.Id, user.Username)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err)
		return
	}
	err = h.TokensRepo.Add(token, expires)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	loginResponseBody := &LoginResponseBody{
		Token: token,
	}
	utils.JSONResponse(w, http.StatusOK, loginResponseBody)
}
