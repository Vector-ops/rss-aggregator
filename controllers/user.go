package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vector-ops/rss-aggregator/internal/database"
	"github.com/Vector-ops/rss-aggregator/models"
	"github.com/Vector-ops/rss-aggregator/utils"
	"github.com/google/uuid"
)

type UserHandler struct {
	store *database.Queries
}

func NewUserHandler(store *database.Queries) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := h.store.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, models.TransformUser(user))
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.TransformUser(user))
}

func (h *UserHandler) GetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := h.store.GetPostsForUsers(r.Context(), database.GetPostsForUsersParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.TransformManyPosts(posts))
}
