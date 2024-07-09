package auth

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Pineapple217/Sortify/web/pkg/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

type Auth struct {
	SessionID   int64
	UserID      int64
	Username    string
	LoggedIn    bool
	SpotifyAuth oauth2.Token
}

func (auth Auth) Check() bool {
	return auth.LoggedIn
}

func (auth Auth) CheckSpotify() bool {
	return auth.SpotifyAuth.Valid()
}

func (auth *Auth) GetClient(
	ctx context.Context,
	spotifyAuth *spotifyauth.Authenticator,
	db *database.Queries,
) *spotify.Client {
	tok, err := spotifyAuth.RefreshToken(ctx, &auth.SpotifyAuth)
	if err != nil {
		panic(err)
	}
	if tok.AccessToken != auth.SpotifyAuth.AccessToken {
		slog.Debug("updating spotify token", "user", auth.UserID)
		err = db.SessionUpdateSotify(ctx, database.SessionUpdateSotifyParams{
			ID:                  auth.SessionID,
			SpotifyAccessToken:  pgtype.Text{String: tok.AccessToken, Valid: true},
			SpotifyRefreshToken: pgtype.Text{String: tok.RefreshToken, Valid: true},
			SpotifyExpiry:       pgtype.Timestamptz{Time: tok.Expiry, Valid: true},
		})
		if err != nil {
			panic(err)
		}
		auth.SpotifyAuth = *tok
	}
	client := spotify.New(spotifyAuth.Client(ctx, tok))
	return client
}

type contextKey string

var AuthContextKey contextKey = "auth"

func AuthMiddleware(db *database.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			a := Auth{
				LoggedIn:    false,
				SpotifyAuth: oauth2.Token{},
			}
			sess, err := session.Get("session", c)
			if err != nil {
				return err
			}
			token, ok := sess.Values["sessionToken"]
			if ok {
				authInfo, err := db.GetAuthInfo(c.Request().Context(), token.(string))
				if err == nil {
					a.LoggedIn = true
					a.UserID = authInfo.ID
					a.Username = authInfo.Username
					if authInfo.SpotifyAccessToken.Valid {
						tok := oauth2.Token{
							AccessToken:  authInfo.SpotifyAccessToken.String,
							RefreshToken: authInfo.SpotifyRefreshToken.String,
							Expiry:       authInfo.SpotifyExpiry.Time,
						}
						a.SpotifyAuth = tok
					}
					a.SessionID = authInfo.ID_2
				}
			}
			ctx := context.WithValue(c.Request().Context(), AuthContextKey, a)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func CheckAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := GetAuth(c.Request().Context())
		slog.Info("testing", "auth", auth.Check())
		if !auth.Check() {
			if len(c.Request().Header.Get("HX-Request")) > 0 {
				c.Response().Header().Set("HX-Redirect", "/login")
				return c.NoContent(http.StatusUnauthorized)
			}
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}

func CheckSpotifyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := GetAuth(c.Request().Context())
		if !auth.Check() {
			if len(c.Request().Header.Get("HX-Request")) > 0 {
				c.Response().Header().Set("HX-Redirect", "/login")
				return c.NoContent(http.StatusUnauthorized)
			}
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		if !auth.CheckSpotify() {
			if len(c.Request().Header.Get("HX-Request")) > 0 {
				c.Response().Header().Set("HX-Redirect", "/spotify_auth")
				return c.NoContent(http.StatusUnauthorized)
			}
			return c.Redirect(http.StatusSeeOther, "/spotify_auth")
		}
		return next(c)
	}
}

func CheckNotAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := GetAuth(c.Request().Context())
		if auth.Check() {
			if len(c.Request().Header.Get("HX-Request")) > 0 {
				c.Response().Header().Set("HX-Redirect", "/")
				return c.NoContent(http.StatusSeeOther)
			}
			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}

func GetAuth(ctx context.Context) Auth {
	if auth, ok := ctx.Value(AuthContextKey).(Auth); ok {
		return auth
	}
	return Auth{LoggedIn: false}
}
