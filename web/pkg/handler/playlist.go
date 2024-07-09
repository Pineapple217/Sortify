package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/database"
	"github.com/Pineapple217/Sortify/web/pkg/util"
	"github.com/Pineapple217/Sortify/web/pkg/view"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func (h *Handler) PlaylistsIndex(c echo.Context) error {
	a := auth.GetAuth(c.Request().Context())
	pls, err := h.DB.GetPlaylistsByUser(c.Request().Context(), a.UserID)
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
	playlist, err := h.DB.GetPlaylist(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return render(c, view.Playlist(playlist))
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

	tracks, err := h.DB.GetPlaylistTracks(c.Request().Context(), database.GetPlaylistTracksParams{
		ID:     id,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
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

	playlist, err := h.DB.CreatePlaylist(c.Request().Context(), database.CreatePlaylistParams{
		Title:  util.RandomString(8),
		UserID: a.UserID,
	})
	if err != nil {
		return err
	}

	tracks, err := client.CurrentUsersTracks(c.Request().Context(), spotify.Limit(50))
	if err != nil {
		return err
	}
	slog.Info("pulling platslist", "total_tracks", tracks.Total)
	InsertTracks(c.Request().Context(), h.DB, tracks.Tracks, playlist.ID)
	for page := 1; ; page++ {
		slog.Info("Pulling page", "page", page, "tracks", len(tracks.Tracks))
		err = client.NextPage(c.Request().Context(), tracks)
		if err != nil && err != spotify.ErrNoMorePages {
			return err
		}
		InsertTracks(c.Request().Context(), h.DB, tracks.Tracks, playlist.ID)
		if err == spotify.ErrNoMorePages {
			break
		}
	}

	return c.NoContent(http.StatusOK)
}

func InsertTracks(ctx context.Context, db *database.Queries, tracks []spotify.SavedTrack, playlistId int64) {
	dbTracks := []database.CreateBulkTracksParams{}
	for _, t := range tracks {
		dbTracks = append(dbTracks, TrackToDBTrack(t))
	}
	sdf := db.CreateBulkTracks(ctx, dbTracks)

	dbRel := []database.CreateBulkTrackPlaylistParams{}
	sdf.QueryRow(func(i1 int, id int64, err error) {
		dbRel = append(dbRel, database.CreateBulkTrackPlaylistParams{
			PlaylistID: playlistId,
			TrackID:    id,
		})
	})
	db.CreateBulkTrackPlaylist(ctx, dbRel).Exec(nil)
	// TODO: error checking
}

func TrackToDBTrack(t spotify.SavedTrack) database.CreateBulkTracksParams {
	dbtrack := database.CreateBulkTracksParams{
		Title:       t.Name,
		Artist:      t.Artists[0].Name,
		ReleaseDate: pgtype.Date{Time: t.Album.ReleaseDateTime(), Valid: true},
		SpotifyID:   t.ID.String(),
		DurationMs:  int32(t.TimeDuration().Seconds()),
		Popularity:  int16(t.Popularity),
	}
	if t.FullTrack.PreviewURL != "" {
		dbtrack.PreviewUrl = pgtype.Text{String: t.FullTrack.PreviewURL, Valid: true}
	}
	if len(t.Album.Images) > 0 {
		dbtrack.ImgSmallUrl.Valid = true
		dbtrack.ImgSmallUrl.String = t.Album.Images[2].URL
		dbtrack.ImgMediumUrl.Valid = true
		dbtrack.ImgMediumUrl.String = t.Album.Images[1].URL
		dbtrack.ImgLargeUrl.Valid = true
		dbtrack.ImgLargeUrl.String = t.Album.Images[0].URL
	}
	return dbtrack
}

// func getClient(c echo.Context, db *database.Queries, spotifyAuth *spotifyauth.Authenticator) (*spotify.Client, error) {
// 	a := auth.GetAuth(c.Request().Context())
// 	if !a.Check() {
// 		return nil, redirect(c, http.StatusUnauthorized, "/login")
// 	}
// 	sess, err := session.Get("session", c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	st, ok := sess.Values["sessionToken"].(string)
// 	if !ok {
// 		return nil, errors.New("session token is not a string")
// 	}
// 	session, err := db.GetSessionByToken(c.Request().Context(), st)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !session.SpotifyAccessToken.Valid {
// 		return nil, redirect(c, http.StatusUnauthorized, "/login")
// 	}
// 	tok := oauth2.Token{
// 		AccessToken:  session.SpotifyAccessToken.String,
// 		RefreshToken: session.SpotifyRefreshToken.String,
// 		Expiry:       session.ExpiresAt.Time,
// 	}
// 	if !tok.Valid() {
// 		return nil, redirect(c, http.StatusUnauthorized, "/login")
// 	}
// 	client := spotify.New(spotifyAuth.Client(c.Request().Context(), &tok))
// 	return client, nil
// }
