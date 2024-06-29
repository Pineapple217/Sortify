package auth

import (
	"context"

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

func GetAuth(ctx context.Context) Auth {
	if auth, ok := ctx.Value(AuthContextKey).(Auth); ok {
		return auth
	}
	return Auth{LoggedIn: false}
}
