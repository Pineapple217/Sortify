package handler

import (
	"errors"
	"log/slog"
	"net/http"

	DBSession "github.com/Pineapple217/Sortify/web/ent/session"
	"github.com/Pineapple217/Sortify/web/pkg/util"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SpotifyAuthCallback(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	state, ok := sess.Values["spotifyAuthState"].(string)
	if !ok {
		return errors.New("spotify state is not a string")
	}

	tok, err := h.SpotifyAuth.Token(c.Request().Context(), state, c.Request())
	if err != nil {
		return err
	}
	if st := c.QueryParam("state"); st != state {
		slog.Warn("aaa", "state", c.QueryParam("state"))
		return c.NoContent(http.StatusBadRequest)
	}

	st, ok := sess.Values["sessionToken"].(string)
	if !ok {
		return errors.New("session token is not a string")
	}

	_, err = h.DB.Session.Update().
		Where(DBSession.Token(st)).
		SetSpotifyAccessToken(tok.AccessToken).
		SetSpotifyRefreshToken(tok.RefreshToken).
		SetSpotifyExpiry(tok.Expiry).
		Save(c.Request().Context())
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) SpotifyLoginUrl(c echo.Context) error {
	r := util.RandomString(16)

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["spotifyAuthState"] = r
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return redirect(c, http.StatusOK, h.SpotifyAuth.AuthURL(r))
}
