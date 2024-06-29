package scheduler

import (
	"context"
	"log/slog"

	"github.com/Pineapple217/Sortify/web/pkg/database"
)

func SessionCleanup(ctx context.Context, db *database.Queries) {
	slog.Debug("deleting expired sessions")
	err := db.DeleteOldSessions(ctx)
	if err != nil {
		slog.Warn("failed to delete expired sessions", "error", err)
	}
}
