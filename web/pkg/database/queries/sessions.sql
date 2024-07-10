-- name: SessionAddSotify :exec
UPDATE sessions
    set
    spotify_access_token = $2,
    spotify_refresh_token = $3,
    spotify_expiry = $4
WHERE token = $1;

-- name: SessionUpdateSotify :exec
UPDATE sessions
    set
    spotify_access_token = $2,
    spotify_refresh_token = $3,
    spotify_expiry = $4
WHERE id = $1;

-- name: DeleteOldSessions :exec
DELETE FROM sessions
WHERE expires_at < now();

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
