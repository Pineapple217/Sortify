package handler

import (
	"github.com/Pineapple217/Sortify/web/pkg/database"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Handler struct {
	DB          *database.Queries
	SpotifyAuth *spotifyauth.Authenticator
}

func NewHandler(db *database.Queries, sp *spotifyauth.Authenticator) *Handler {
	return &Handler{
		DB:          db,
		SpotifyAuth: sp,
	}
}
