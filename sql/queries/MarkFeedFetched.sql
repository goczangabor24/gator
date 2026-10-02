-- name: MarkFeedFetched :one

UPDATE feeds
SET last_fetched_at = current_timestamp, updated_at = current_timestamp 
WHERE id = $1
RETURNING *;