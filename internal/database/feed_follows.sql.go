// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollows = `-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at,user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowsParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFeedFollows(ctx context.Context, arg CreateFeedFollowsParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollows,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}
