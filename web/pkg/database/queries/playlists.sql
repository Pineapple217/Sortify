
-- name: GetPlaylist :one
select *
from playlists
where id = $1;

-- name: GetPlaylistsByUser :many
SELECT *
from playlists
where user_id = $1 and deleted_at is null
order by updated_at desc;

-- name: GetPlaylistTracks :many
select t.* from playlists p
join playlist_track pt on pt.playlist_id = p.id
join tracks t on pt.track_id = t.id
where p.id = $1
limit $2
offset $3;

-- name: CreatePlaylist :one
INSERT INTO playlists (
    title, user_id
) VALUES (
    $1, $2
) RETURNING *;