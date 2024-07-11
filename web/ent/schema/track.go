package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Track holds the schema definition for the Track entity.
type Track struct {
	ent.Schema
}

// Fields of the Track.
func (Track) Fields() []ent.Field {
	return []ent.Field{
		// TODO: validation
		field.String("name"),
		field.String("artist"),

		field.String("img_small_url").Optional().Nillable(),
		field.String("img_medium_url").Optional().Nillable(),
		field.String("img_large_url").Optional().Nillable(),

		// TODO: Date type
		field.Time("release_date"),
		field.Text("spotify_id").Unique().Immutable(),
		field.Int("duration_ms").NonNegative(),
		field.String("preview_url").Optional().Nillable(),
		field.Int("popularity"),
		// TODO: validation

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playlists", Playlist.Type).Ref("tracks"),
	}
}
