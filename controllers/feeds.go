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

type FeedHandler struct {
	store *database.Queries
}

func NewFeedHandler(store *database.Queries) *FeedHandler {
	return &FeedHandler{
		store: store,
	}
}

func (h *FeedHandler) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := h.store.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, models.TransformFeed(feed))
}

func (h *FeedHandler) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.store.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.TransformManyFeeds(feeds))
}
