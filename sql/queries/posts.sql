-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsbyUser :many
SELECT p.id, p.created_at, p.updated_at, p.title, p.url, p.description, p.published_at, f.name feed_name
FROM posts p, feeds f, feed_follows ff
WHERE p.feed_id = f.id
  AND ff.followed_feed_id = f.id
  AND ff.following_user_id = $1
ORDER BY p.published_at DESC
;
