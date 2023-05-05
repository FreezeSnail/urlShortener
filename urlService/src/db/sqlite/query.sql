-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name, email
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: GetURL :one
SELECT * FROM urls
WHERE url = ? LIMIT 1;

-- name: ListUrls :many
SELECT * FROM urls
ORDER BY url;

-- name: CreateUrl :one
INSERT INTO urls (
  url, shorturl, userid, createdate
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: DeleteUrl :exec
DELETE FROM urls
WHERE id = ?;