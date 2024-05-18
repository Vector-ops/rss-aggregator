package models

import (
	"database/sql"
	"time"

	"github.com/Vector-ops/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID    `json:"id,omitempty"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	Name      string       `json:"name,omitempty"`
	Url       string       `json:"url,omitempty"`
	UserID    uuid.UUID    `json:"user_id,omitempty"`
	LastFetch sql.NullTime `json:"last_fetch"`
}

func TransformFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func TransformManyFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, Feed(dbFeed))
	}

	return feeds
}
