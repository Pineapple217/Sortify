package auth

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Pineapple217/Sortify/web/ent"
	DBSession "github.com/Pineapple217/Sortify/web/ent/session"
	"github.com/Pineapple217/Sortify/web/ent/user"
	"github.com/Pineapple217/Sortify/web/pkg/util"
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
	db *ent.Client,
) *spotify.Client {
	tok, err := spotifyAuth.RefreshToken(ctx, &auth.SpotifyAuth)
	if err != nil {
		panic(err)
	}
	if tok.AccessToken != auth.SpotifyAuth.AccessToken {
		slog.Debug("updating spotify token", "user", auth.UserID)
		_, err = db.Session.Update().
			Where(DBSession.ID(int(auth.SessionID))).
			SetSpotifyAccessToken(tok.AccessToken).
			SetSpotifyRefreshToken(tok.RefreshToken).
			SetExpiresAt(tok.Expiry).
			Save(ctx)
		util.MaybeDieErr(err)
		auth.SpotifyAuth = *tok
	}
	client := spotify.New(spotifyAuth.Client(ctx, tok))
	return client
}

type contextKey string

var AuthContextKey contextKey = "auth"

func AuthMiddleware(db *ent.Client) echo.MiddlewareFunc {
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
				// TODO: does not use a join!!!
				s, err := db.Session.Query().
					Where(DBSession.Token(token.(string))).
					WithUser(func(uq *ent.UserQuery) {
						uq.Select(user.FieldID, user.FieldUsername)
					}).Only(c.Request().Context())
				if _, ok := err.(*ent.NotFoundError); !ok {
					if err != nil {
						panic(err)
					}
					a.LoggedIn = true
					a.UserID = int64(s.ID)
					a.Username = s.Edges.User.Username
					if s.SpotifyAccessToken != nil {
						tok := oauth2.Token{
							AccessToken:  *s.SpotifyAccessToken,
							RefreshToken: *s.SpotifyRefreshToken,
							Expiry:       s.ExpiresAt,
						}
						a.SpotifyAuth = tok
					}
					a.SessionID = int64(s.ID)
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
