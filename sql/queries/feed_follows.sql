-- name: CreateFeedFollow :one
WITH inserted_record AS (
    INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4)
    RETURNING *
)
SELECT
inserted_record.*,
feeds.name AS feed_name,
users.name AS user_name
FROM inserted_record
INNER JOIN feeds
ON feeds.id = inserted_record.feed_id
INNER JOIN users
ON users.id = inserted_record.user_id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, users.name, feeds.name as feed_name FROM feed_follows
INNER JOIN users
ON feed_follows.user_id = users.id
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;