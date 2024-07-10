-- name: CreateBulkTrackPlaylist :batchexec
INSERT INTO playlist_track (
    playlist_id, track_id
) VALUES (
    $1, $2
);

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
