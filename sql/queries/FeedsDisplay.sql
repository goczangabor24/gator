-- name: FeedsDisplay :many

SELECT feeds.name, feeds.url, users.name
FROM feeds
JOIN users ON users.id = feeds.user_id;