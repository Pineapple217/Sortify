-- name: CreateUser :one
INSERT INTO users (
    username, password_hash
) VALUES (
    $1, $2
) RETURNING *;

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE token = $1;

-- name: GetSessionByToken :one
SELECT * FROM sessions
WHERE token = $1;

-- name: CreateSession :one
INSERT INTO sessions (
    user_id, token, expires_at, ip_address, user_agent
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserBySession :one
SELECT u.* FROM users u
JOIN sessions s on s.user_id = u.id
WHERE s.token = $1
AND expires_at > now();

-- name: SessionAddSotify :exec
UPDATE sessions
    set
    spotify_access_token = $2,
    spotify_refresh_token = $3,
    spotify_expiry = $4
WHERE token = $1;


-- name: DeleteOldSessions :exec
DELETE FROM sessions
WHERE expires_at < now();