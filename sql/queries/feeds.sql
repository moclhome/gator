-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, created_by_user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.id, f.created_at, f.updated_at, f.name feed_name, f.url, u.name user_name
FROM feeds f, users u
WHERE f.created_by_user_id = u.id
;

-- name: GetFeed :one
SELECT f.id, f.created_at, f.updated_at, f.name, f.url
FROM feeds f
WHERE f.url = $1
;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_DATE,
    updated_at = CURRENT_DATE
WHERE id = $1
;

-- name: GetNextFeedToFetch :one
SELECT f.*
FROM feeds f, feed_follows ff
WHERE f.id = ff.followed_feed_id
AND ff.following_user_id = $1
ORDER BY f.last_fetched_at NULLS FIRST
LIMIT 1
;