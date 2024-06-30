package auth

import (
	"context"
	"net/http"

	"github.com/Pineapple217/Sortify/web/pkg/database"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	UserID      int64
	Username    string
	LoggedIn    bool
	SpotifyAuth bool
}

func (auth Auth) Check() bool {
	return auth.LoggedIn
}

func (auth Auth) CheckSpotify() bool {
	return auth.SpotifyAuth
}

type contextKey string

var AuthContextKey contextKey = "auth"

func AuthMiddleware(db *database.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			a := Auth{
				LoggedIn:    false,
				SpotifyAuth: false,
			}
			sess, err := session.Get("session", c)
			if err != nil {
				return err
			}
			token, ok := sess.Values["sessionToken"]
			if ok {
				user, err := db.GetUserBySession(c.Request().Context(), token.(string))
				if err == nil {
					a.LoggedIn = true
					a.UserID = user.ID
					a.Username = user.Username
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
		if !auth.Check() {
			if len(c.Request().Header.Get("HX-Request")) > 0 {
				c.Response().Header().Set("HX-Redirect", "/login")
				return c.NoContent(http.StatusUnauthorized)
			}
			return c.Redirect(http.StatusUnauthorized, "/login")
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
			return c.Redirect(http.StatusUnauthorized, "/login")
		}
		if !auth.CheckSpotify() {
			if len(c.Request().Header.Get("HX-Request")) > 0 {
				c.Response().Header().Set("HX-Redirect", "/spotify_auth")
				return c.NoContent(http.StatusUnauthorized)
			}
			return c.Redirect(http.StatusUnauthorized, "/spotify_auth")
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
