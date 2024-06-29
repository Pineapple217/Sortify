// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    user_id, token, expires_at, ip_address, user_agent
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, user_id, token, ip_address, user_agent, expires_at, last_login_at, created_at, spotify_access_token, spotify_refresh_token, spotify_expiry
`

type CreateSessionParams struct {
	UserID    int64
	Token     string
	ExpiresAt pgtype.Timestamptz
	IpAddress pgtype.Text
	UserAgent pgtype.Text
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.UserID,
		arg.Token,
		arg.ExpiresAt,
		arg.IpAddress,
		arg.UserAgent,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.IpAddress,
		&i.UserAgent,
		&i.ExpiresAt,
		&i.LastLoginAt,
		&i.CreatedAt,
		&i.SpotifyAccessToken,
		&i.SpotifyRefreshToken,
		&i.SpotifyExpiry,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, password_hash
) VALUES (
    $1, $2
) RETURNING id, username, password_hash, created_at, updated_at
`

type CreateUserParams struct {
	Username     string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOldSessions = `-- name: DeleteOldSessions :exec
DELETE FROM sessions
WHERE expires_at < now()
`

func (q *Queries) DeleteOldSessions(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteOldSessions)
	return err
}

const deleteSessionByToken = `-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE token = $1
`

func (q *Queries) DeleteSessionByToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteSessionByToken, token)
	return err
}

const getSessionByToken = `-- name: GetSessionByToken :one
SELECT id, user_id, token, ip_address, user_agent, expires_at, last_login_at, created_at, spotify_access_token, spotify_refresh_token, spotify_expiry FROM sessions
WHERE token = $1
`

func (q *Queries) GetSessionByToken(ctx context.Context, token string) (Session, error) {
	row := q.db.QueryRow(ctx, getSessionByToken, token)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.IpAddress,
		&i.UserAgent,
		&i.ExpiresAt,
		&i.LastLoginAt,
		&i.CreatedAt,
		&i.SpotifyAccessToken,
		&i.SpotifyRefreshToken,
		&i.SpotifyExpiry,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password_hash, created_at, updated_at FROM users
WHERE username = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserBySession = `-- name: GetUserBySession :one
SELECT u.id, u.username, u.password_hash, u.created_at, u.updated_at FROM users u
JOIN sessions s on s.user_id = u.id
WHERE s.token = $1
AND expires_at > now()
`

func (q *Queries) GetUserBySession(ctx context.Context, token string) (User, error) {
	row := q.db.QueryRow(ctx, getUserBySession, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const sessionAddSotify = `-- name: SessionAddSotify :exec
UPDATE sessions
    set
    spotify_access_token = $2,
    spotify_refresh_token = $3,
    spotify_expiry = $4
WHERE token = $1
`

type SessionAddSotifyParams struct {
	Token               string
	SpotifyAccessToken  pgtype.Text
	SpotifyRefreshToken pgtype.Text
	SpotifyExpiry       pgtype.Timestamptz
}

func (q *Queries) SessionAddSotify(ctx context.Context, arg SessionAddSotifyParams) error {
	_, err := q.db.Exec(ctx, sessionAddSotify,
		arg.Token,
		arg.SpotifyAccessToken,
		arg.SpotifyRefreshToken,
		arg.SpotifyExpiry,
	)
	return err
}
