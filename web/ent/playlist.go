// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Pineapple217/Sortify/web/ent/playlist"
	"github.com/Pineapple217/Sortify/web/ent/user"
)

// Playlist is the model entity for the Playlist schema.
type Playlist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaylistQuery when eager-loading is set.
	Edges          PlaylistEdges `json:"edges"`
	user_playlists *int
	selectValues   sql.SelectValues
}

// PlaylistEdges holds the relations/edges for other nodes in the graph.
type PlaylistEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Tracks holds the value of the tracks edge.
	Tracks []*Track `json:"tracks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaylistEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TracksOrErr returns the Tracks value or an error if the edge
// was not loaded in eager-loading.
func (e PlaylistEdges) TracksOrErr() ([]*Track, error) {
	if e.loadedTypes[1] {
		return e.Tracks, nil
	}
	return nil, &NotLoadedError{edge: "tracks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Playlist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case playlist.FieldID:
			values[i] = new(sql.NullInt64)
		case playlist.FieldName:
			values[i] = new(sql.NullString)
		case playlist.FieldDeletedAt, playlist.FieldCreatedAt, playlist.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case playlist.ForeignKeys[0]: // user_playlists
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Playlist fields.
func (pl *Playlist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playlist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case playlist.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pl.DeletedAt = value.Time
			}
		case playlist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case playlist.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		case playlist.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case playlist.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_playlists", value)
			} else if value.Valid {
				pl.user_playlists = new(int)
				*pl.user_playlists = int(value.Int64)
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Playlist.
// This includes values selected through modifiers, order, etc.
func (pl *Playlist) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Playlist entity.
func (pl *Playlist) QueryUser() *UserQuery {
	return NewPlaylistClient(pl.config).QueryUser(pl)
}

// QueryTracks queries the "tracks" edge of the Playlist entity.
func (pl *Playlist) QueryTracks() *TrackQuery {
	return NewPlaylistClient(pl.config).QueryTracks(pl)
}

// Update returns a builder for updating this Playlist.
// Note that you need to call Playlist.Unwrap() before calling this method if this Playlist
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Playlist) Update() *PlaylistUpdateOne {
	return NewPlaylistClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Playlist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Playlist) Unwrap() *Playlist {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Playlist is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Playlist) String() string {
	var builder strings.Builder
	builder.WriteString("Playlist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("deleted_at=")
	builder.WriteString(pl.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Playlists is a parsable slice of Playlist.
type Playlists []*Playlist
