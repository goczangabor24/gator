-- name: Unfollow :exec

WITH feedToDelete AS (
    SELECT feed_follows.id
    FROM feed_follows
    JOIN feeds ON feed_follows.feed_id = feeds.id
    JOIN users ON feed_follows.user_id = users.id
    WHERE feeds.url = $1 AND users.id = $2
)

DELETE FROM feed_follows
WHERE id IN (SELECT id FROM feedToDelete);