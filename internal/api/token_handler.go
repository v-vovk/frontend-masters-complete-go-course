package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/v-vovk/femProject/internal/store"
	"github.com/v-vovk/femProject/internal/tokens"
	"github.com/v-vovk/femProject/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

func (h *TokenHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		h.logger.Printf("ERROR: createTokenRequest: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request payload"})
		return
	}

	user, err := h.userStore.GetByUsername(req.Username)
	if err != nil || user == nil {
		h.logger.Printf("ERROR: GetByUsername: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	passwordsMatching, err := user.PasswordHash.Matches(req.Password)
	if err != nil {
		h.logger.Printf("ERORR: PasswordHash.Mathes %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if !passwordsMatching {
		h.logger.Printf("ERROR: password mismatch for user %s", req.Username)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid credentials"})

		return
	}

	token, err := h.tokenStore.CreateToken(user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR: creating token: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})

		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"auth_token": token})

}
