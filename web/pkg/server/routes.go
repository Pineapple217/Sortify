package server

import (
	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/handler"
	"github.com/Pineapple217/Sortify/web/pkg/static"
	"github.com/labstack/echo/v4"
)

func (server *Server) RegisterRoutes(hdlr *handler.Handler) {
	e := server.e

	s := e.Group("/static")
	s.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("Cache-Control", "public, max-age=31536000, immutable")
			return next(c)
		}
	})
	s.StaticFS("/", echo.MustSubFS(static.PublicFS, "public"))

	IsAuth := e.Group("", auth.CheckAuthMiddleware)
	IsSpotAuth := e.Group("", auth.CheckSpotifyAuthMiddleware)
	_ = IsSpotAuth

	e.GET("/", hdlr.Home)
	IsAuth.GET("/callback", hdlr.SpotifyAuthCallback)
	IsAuth.GET("/spotify_auth", hdlr.SpotifyLoginUrl)

	noAuth := e.Group("", auth.CheckNotAuthMiddleware)

	noAuth.GET("/login", hdlr.LoginIndex)
	noAuth.POST("/login", hdlr.LoginUser)
	IsAuth.DELETE("/logout", hdlr.LogoutUser)

	noAuth.GET("/signup", hdlr.SignupForm)
	noAuth.POST("/signup", hdlr.SignupUser)

	IsSpotAuth.GET("/playlist/pull", hdlr.PullLiked)
	IsAuth.GET("/playlist", hdlr.PlaylistsIndex)

	IsAuth.GET("/playlist/:id", hdlr.PlaylistIndex)
	IsAuth.GET("/playlist/:id/tracks", hdlr.PlaylistTracks)
}
