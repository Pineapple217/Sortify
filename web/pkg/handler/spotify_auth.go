package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/database"
	"github.com/Pineapple217/Sortify/web/pkg/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SpotifyAuthCallback(c echo.Context) error {
	auth := auth.GetAuth(c.Request().Context())
	if !auth.Check() {
		return c.NoContent(http.StatusForbidden)
	}

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

	err = h.DB.SessionAddSotify(c.Request().Context(), database.SessionAddSotifyParams{
		Token:               st,
		SpotifyAccessToken:  pgtype.Text{String: tok.AccessToken, Valid: true},
		SpotifyRefreshToken: pgtype.Text{String: tok.AccessToken, Valid: true},
		SpotifyExpiry:       pgtype.Timestamptz{Time: tok.Expiry, Valid: true},
	})
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) SpotifyLoginUrl(c echo.Context) error {
	auth := auth.GetAuth(c.Request().Context())
	if !auth.Check() {
		return c.NoContent(http.StatusForbidden)
	}

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

	c.Response().Header().Set("HX-Redirect", h.SpotifyAuth.AuthURL(r))
	return c.NoContent(http.StatusOK)
}
