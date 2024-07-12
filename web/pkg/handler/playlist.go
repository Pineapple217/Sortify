package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Pineapple217/Sortify/web/ent"
	"github.com/Pineapple217/Sortify/web/ent/playlist"
	"github.com/Pineapple217/Sortify/web/ent/user"
	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/view"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func (h *Handler) PlaylistsIndex(c echo.Context) error {
	a := auth.GetAuth(c.Request().Context())
	pls, err := h.DB.User.Query().
		Where(user.ID(int(a.UserID))).QueryPlaylists().
		Order(playlist.ByCreatedAt()).
		All(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.PlaylistIndex(pls))
}

func (h *Handler) PlaylistIndex(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	playlist, err := h.DB.Playlist.Get(c.Request().Context(), int(id))
	if err != nil {
		return err
	}
	return render(c, view.Playlist(playlist))
}

func (h *Handler) PlaylistDelete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = h.DB.Playlist.DeleteOneID(int(id)).Exec(c.Request().Context())
	if ent.IsNotFound(err) {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}

func (h *Handler) PlaylistTracks(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	offsetStr := c.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tracks, err := h.DB.Playlist.Query().
		Where(playlist.ID(int(id))).
		QueryTracks().
		Limit(limit).
		Offset(offset).
		All(c.Request().Context())
	slog.Info("fetched tracks", "count", len(tracks))
	if err != nil {
		return err
	}
	return render(c, view.PlaylistTracks(view.PlaylistTracksData{
		Tracks: tracks,
		Offset: offset,
		Limit:  limit,
	}))
}

func (h *Handler) PullLiked(c echo.Context) error {
	a := auth.GetAuth(c.Request().Context())
	client := a.GetClient(c.Request().Context(), h.SpotifyAuth, h.DB)

	playlist, err := h.DB.Playlist.Create().
		SetName("Liked - " + time.Now().Format("2006-01-02")).
		SetUserID(int(a.UserID)).
		Save(c.Request().Context())

	if err != nil {
		return err
	}

	tracks, err := client.CurrentUsersTracks(c.Request().Context(), spotify.Limit(50))
	if err != nil {
		return err
	}
	slog.Info("pulling platslist", "total_tracks", tracks.Total)
	InsertTracks(c.Request().Context(), h.DB, tracks.Tracks, int64(playlist.ID))
	for page := 1; ; page++ {
		slog.Info("Pulling page", "page", page, "tracks", len(tracks.Tracks))
		err = client.NextPage(c.Request().Context(), tracks)
		if err != nil && err != spotify.ErrNoMorePages {
			return err
		}
		InsertTracks(c.Request().Context(), h.DB, tracks.Tracks, int64(playlist.ID))
		if err == spotify.ErrNoMorePages {
			break
		}
	}

	return c.NoContent(http.StatusOK)
}

func InsertTracks(ctx context.Context, db *ent.Client, tracks []spotify.SavedTrack, playlistId int64) {
	err := db.Track.MapCreateBulk(tracks, func(c *ent.TrackCreate, i int) {
		c.SetName(tracks[i].Name)
		c.SetArtist(tracks[i].Artists[0].Name)
		c.SetReleaseDate(tracks[i].Album.ReleaseDateTime())
		c.SetSpotifyID(tracks[i].ID.String())
		c.SetDurationMs(int(tracks[i].TimeDuration().Microseconds()))
		c.SetPopularity(int(tracks[i].Popularity))
		if tracks[i].PreviewURL != "" {
			c.SetPreviewURL(tracks[i].PreviewURL)
		}
		if len(tracks[i].Album.Images) > 0 {
			c.SetImgSmallURL(tracks[i].Album.Images[2].URL)
			c.SetImgMediumURL(tracks[i].Album.Images[1].URL)
			c.SetImgLargeURL(tracks[i].Album.Images[0].URL)
		}
		c.AddPlaylistIDs(int(playlistId))
	}).OnConflict(
		sql.ConflictColumns("spotify_id"),
	).UpdateNewValues().Exec(ctx)
	if err != nil {
		panic(err)
	}
}
