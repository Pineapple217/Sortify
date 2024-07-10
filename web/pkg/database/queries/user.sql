
-- name: GetUser :one
SELECT * FROM users
WHERE username = $1;

-- name: GetAuthInfo :one
SELECT u.username, u.id,
    s.id, s.spotify_access_token, s.spotify_refresh_token, s.spotify_expiry
FROM users u
JOIN sessions s on s.user_id = u.id
WHERE s.token = $1
AND expires_at > now();

-- name: CreateUser :one
INSERT INTO users (
    username, password_hash
) VALUES (
    $1, $2
) RETURNING *;
