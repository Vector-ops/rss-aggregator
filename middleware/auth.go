package middleware

import (
	"fmt"
	"net/http"

	"github.com/Vector-ops/rss-aggregator/internal/auth"
	"github.com/Vector-ops/rss-aggregator/internal/database"
	"github.com/Vector-ops/rss-aggregator/utils"
)

type AuthHandlerFunc func(http.ResponseWriter, *http.Request, database.User)

type AuthHandler struct {
	store *database.Queries
}

func NewAuthHandler(store *database.Queries) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) MiddlewareAuth(handler AuthHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		}

		user, err := h.store.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
		}
		handler(w, r, user)
	}
}
