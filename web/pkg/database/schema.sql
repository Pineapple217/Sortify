CREATE TABLE IF NOT EXISTS users (
    id                BIGSERIAL PRIMARY KEY,
    username          text      NOT NULL UNIQUE,
    password_hash     text      NOT NULL, 
    -- TODO: salt text NOT NULL,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
    id                BIGSERIAL PRIMARY KEY,
    user_id           BIGINT REFERENCES users (id) NOT NULL,
    token             text      NOT NULL, 
    ip_address        text, 
    user_agent        text, 
    expires_at        TIMESTAMP WITH TIME ZONE NOT NULL,
    last_login_at     TIMESTAMP WITH TIME ZONE,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    spotify_access_token    text,
    spotify_refresh_token   text,
    spotify_expiry          TIMESTAMP WITH TIME ZONE
);

-- CREATE TABLE IF NOT EXISTS playlists (
--     id                BIGSERIAL PRIMARY KEY,
--     name              text NOT NULL,
--     -- TODO public bool
--     user_id           BIGINT REFERENCES users (id) NOT NULL,

--     created_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
--     updated_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
--     deleted_at        TIMESTAMP WITH TIME ZONE DEFAULT
-- );

-- CREATE TABLE IF NOT EXISTS playlist_track (
--     id                BIGSERIAL PRIMARY KEY,
--     playlist_id       BIGINT REFERENCES playlists (id) NOT NULL,
--     tack_id           BIGINT REFERENCES tracks (id) NOT NULL,

--     added_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
-- );

-- CREATE TABLE IF NOT EXISTS tracks (
--     id                BIGSERIAL PRIMARY KEY,
--     name              text NOT NULL,
--     artist            text NOT NULL,

--     img_small_url     text NOT NULL,
--     img_medium_url    text NOT NULL,
--     img_large_url     text NOT NULL,

--     release_date      DATE NOT NULL,

--     spotify_id        text NOT NULL,
--     duration_ms       integer NOT NULL,
--     preview_url       text,
--     popularity        smallint NOT NULL
-- );
