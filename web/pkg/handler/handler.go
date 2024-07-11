package handler

import (
	"github.com/Pineapple217/Sortify/web/ent"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Handler struct {
	DB          *ent.Client
	SpotifyAuth *spotifyauth.Authenticator
}

func NewHandler(client *ent.Client, sp *spotifyauth.Authenticator) *Handler {
	return &Handler{
		DB:          client,
		SpotifyAuth: sp,
	}
}
