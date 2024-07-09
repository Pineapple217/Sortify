package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/Pineapple217/Sortify/web/pkg/config"
	"github.com/Pineapple217/Sortify/web/pkg/database"
	"github.com/Pineapple217/Sortify/web/pkg/handler"
	"github.com/Pineapple217/Sortify/web/pkg/logger"
	"github.com/Pineapple217/Sortify/web/pkg/scheduler"
	"github.com/Pineapple217/Sortify/web/pkg/server"
	"github.com/Pineapple217/Sortify/web/pkg/static"
	"github.com/Pineapple217/Sortify/web/pkg/util"

	_ "github.com/joho/godotenv/autoload"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const banner = `
 _______  _____   ______ _______ _____ _______ __   __
 |______ |     | |_____/    |      |   |______   \_/  
 ______| |_____| |    \_    |    __|__ |          |  

> https://github.com/Pineapple217/Sortify
===- v 0.0.0 -=========================================

`

const redirectURI = "http://localhost:3000/callback"

func main() {
	fmt.Print(banner)

	cnf, err := config.Load()
	util.MaybeDie(err, "Failed to load configs")

	slog.SetDefault(logger.NewLogger(cnf.Logger))

	rr := static.HashPublicFS()

	db := database.NewQueries(cnf.Database)

	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopeUserLibraryRead,
		),
	)

	h := handler.NewHandler(db, auth)

	server := server.NewServer(cnf.Server)
	server.RegisterRoutes(h)
	server.ApplyMiddleware(cnf, rr, db)
	server.Start()
	defer server.Stop()

	s := scheduler.NewScheduler()
	s.Schedule(time.Minute*5, func() {
		scheduler.SessionCleanup(context.Background(), db)
	})
	defer s.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received an interrupt signal, exiting...")
}
