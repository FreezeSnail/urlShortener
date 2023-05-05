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

-- name: ListURLs :many
SELECT * FROM urls
ORDER BY url;

-- name: CreateURL :one
INSERT INTO urls (
  url, shorturl, userid, createdate
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: GetShortURLFromLong :one
SELECT shorturl FROM urls
WHERE url = ? LIMIT 1;

-- name: GetLongURLFromShort :one
SELECT url FROM urls
WHERE shorturl = ? LIMIT 1;

-- name: DeleteUrl :exec
DELETE FROM urls
WHERE id = ?;