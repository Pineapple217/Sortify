package server

import (
	"log/slog"

	"github.com/Pineapple217/Sortify/web/ent"
	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func (s *Server) ApplyMiddleware(cnf config.Configs, reRoutes map[string]string, db *ent.Client) {
	s.e.Pre(echoMw.Rewrite(reRoutes))
	s.e.Use(echoMw.RequestLoggerWithConfig(echoMw.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v echoMw.RequestLoggerValues) error {
			slog.Info("request",
				"method", v.Method,
				"status", v.Status,
				"latency", v.Latency,
				"path", v.URI,
			)
			return nil

		},
	}))
	s.e.Use(echoMw.GzipWithConfig(echoMw.GzipConfig{
		Level: 5,
	}))
	s.e.Use(session.Middleware(sessions.NewCookieStore([]byte(cnf.Server.Secret))))
	s.e.Use(auth.AuthMiddleware(db))
}
