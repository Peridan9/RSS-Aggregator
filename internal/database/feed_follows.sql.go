// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING id, created_at, updated_at, feed_id, user_id
)
SELECT
    inserted.id, inserted.created_at, inserted.updated_at, inserted.feed_id, inserted.user_id,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted
INNER JOIN feeds ON inserted.feed_id = feeds.id
INNER JOIN users ON inserted.user_id = users.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
		&i.UserID,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.id, feed_follows.created_at, feed_follows.updated_at, feed_follows.feed_id, feed_follows.user_id,
    feeds.name AS feed_name,
    feeds.url AS feed_url
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id=$1
`

type GetFeedFollowsForUserRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
	FeedName  string
	FeedUrl   string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
			&i.FeedName,
			&i.FeedUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
