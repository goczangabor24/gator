-- name: DeleteUser :one

DELETE FROM users
WHERE name = $1
RETURNING *;