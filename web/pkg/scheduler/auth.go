package scheduler

import (
	"context"
	"log/slog"
	"time"

	"github.com/Pineapple217/Sortify/web/ent"
	"github.com/Pineapple217/Sortify/web/ent/session"
)

func SessionCleanup(ctx context.Context, db *ent.Client) {
	slog.Debug("deleting expired sessions")
	_, err := db.Session.Delete().
		Where(session.ExpiresAtLT(time.Now())).
		Exec(ctx)
	if err != nil {
		slog.Warn("failed to delete expired sessions", "error", err)
	}
}
