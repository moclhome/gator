-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, following_user_id, followed_feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id, created_at, updated_at,
    (SELECT u.name FROM users u
    WHERE u.id = following_user_id) user_name,
    (SELECT f.name FROM feeds f
    WHERE f.id = followed_feed_id) feed_name
;

-- name: GetFeedFollowsForUser :many
SELECT ff.id, ff.created_at, ff.updated_at, u.name user_name, f.name feed_name
FROM feed_follows ff, users u, feeds f
WHERE ff.following_user_id = u.id
  AND ff.followed_feed_id = f.id
  AND u.name = $1
;

-- name: DeleteFeedFollow :exec
DELETE
FROM feed_follows ff
WHERE ff.following_user_id = $1
  AND ff.followed_feed_id = $2
;
