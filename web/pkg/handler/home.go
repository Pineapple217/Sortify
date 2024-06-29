package handler

import (
	"errors"

	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/view"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

func (h *Handler) Home(c echo.Context) error {
	a := auth.GetAuth(c.Request().Context())
	if !a.Check() {
		return render(c, view.Home("a"))
	}
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	st, ok := sess.Values["sessionToken"].(string)
	if !ok {
		return errors.New("session token is not a string")
	}
	session, err := h.DB.GetSessionByToken(c.Request().Context(), st)
	if err != nil {
		return render(c, view.Home("a"))
	}
	if !session.SpotifyAccessToken.Valid {
		return render(c, view.Home("a"))
	}
	tok := oauth2.Token{
		AccessToken:  session.SpotifyAccessToken.String,
		RefreshToken: session.SpotifyRefreshToken.String,
		Expiry:       session.ExpiresAt.Time,
	}
	client := spotify.New(h.SpotifyAuth.Client(c.Request().Context(), &tok))
	user, err := client.CurrentUser(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.Home(user.DisplayName))
}
