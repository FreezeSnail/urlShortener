-- name: CreateUser :exec
INSERT INTO users (
  name, password, apikey
) VALUES (
  ?, ?, ?
);

-- name: GetAPIKey :one
SELECT apikey FROM users
WHERE name=? AND password=?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: GetHashPassword :one
SELECT password FROM users
where name=?;

-- name: APIKeyValid :one
SELECT COUNT(DISTINCT apikey) FROM users 
WHERE apikey=?;


-- name: GetURL :one
SELECT shorturl FROM urls
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