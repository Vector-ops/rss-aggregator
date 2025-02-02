package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vector-ops/rss-aggregator/internal/database"
	"github.com/Vector-ops/rss-aggregator/models"
	"github.com/Vector-ops/rss-aggregator/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FeedFollowHandler struct {
	store *database.Queries
}

func NewFeedFollowHandler(store *database.Queries) *FeedFollowHandler {
	return &FeedFollowHandler{
		store: store,
	}
}

func (h *FeedFollowHandler) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := h.store.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follows: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, models.TransformFeedFollow(feedFollow))
}

func (h *FeedFollowHandler) GetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := h.store.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follows: %v", err))
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, models.TransformManyFeedFollows(feedFollows))
}

func (h *FeedFollowHandler) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}
	err = h.store.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
