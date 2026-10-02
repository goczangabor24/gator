-- name: GetPostsForUser :many

SELECT title, description
FROM posts
ORDER BY updated_at
LIMIT $1;