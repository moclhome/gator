-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    following_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    followed_feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    constraint feed_follows_uq unique (following_user_id, followed_feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
