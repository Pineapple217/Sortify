package database

import (
	"context"
	"fmt"
	"log/slog"

	"entgo.io/ent/dialect"
	_ "github.com/Pineapple217/Sortify/web/ent/runtime"

	"github.com/Pineapple217/Sortify/web/ent"
	"github.com/Pineapple217/Sortify/web/pkg/config"
	"github.com/Pineapple217/Sortify/web/pkg/util"
	_ "github.com/lib/pq"
)

func NewClient(cnf config.Database) *ent.Client {
	ctx := context.Background()

	DbUser := cnf.User
	DbPassword := cnf.Paddword
	DbDatabase := cnf.Database
	DbHost := cnf.Host
	DbPort := cnf.Port

	slog.Info("Starting database", "host", DbHost, "database", DbDatabase)
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		DbUser, DbPassword, DbDatabase, DbHost, DbPort)

	client, err := ent.Open(dialect.Postgres, connStr)
	util.MaybeDie(err, "Failed to connected to database")

	err = client.Schema.Create(ctx)
	util.MaybeDie(err, "Failed to run migrations")

	return client
}
