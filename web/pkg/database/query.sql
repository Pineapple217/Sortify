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

-- name: GetAuthInfo :one
SELECT u.username, u.id,
    s.id, s.spotify_access_token, s.spotify_refresh_token, s.spotify_expiry
FROM users u
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

-- name: CreatePlaylist :one
INSERT INTO playlists (
    title, user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: CreateBulkTracks :batchone
INSERT INTO tracks (
    title, artist, img_small_url, img_medium_url, img_large_url,
    release_date, spotify_id, duration_ms, preview_url, popularity
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) ON CONFLICT (spotify_id) DO UPDATE SET
    title = EXCLUDED.title,
    artist = EXCLUDED.artist,
    img_small_url = EXCLUDED.img_small_url,
    img_medium_url = EXCLUDED.img_medium_url,
    img_large_url = EXCLUDED.img_large_url,
    release_date = EXCLUDED.release_date,
    duration_ms = EXCLUDED.duration_ms,
    preview_url = EXCLUDED.preview_url,
    popularity = EXCLUDED.popularity,
    updated_at = NOW()
RETURNING id;

-- name: CreateBulkTrackPlaylist :batchexec
INSERT INTO playlist_track (
    playlist_id, track_id
) VALUES (
    $1, $2
);

-- name: GetPlaylistTracks :many
select t.* from playlists p
join playlist_track pt on pt.playlist_id = p.id
join tracks t on pt.track_id = t.id
where p.id = $1
limit $2
offset $3;

-- name: GetPlaylist :one
select *
from playlists
where id = $1;

-- name: GetPlaylistsByUser :many
SELECT *
from playlists
where user_id = $1
order by updated_at desc;
