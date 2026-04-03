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