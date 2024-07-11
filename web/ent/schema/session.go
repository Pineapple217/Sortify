package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		// TODO: make UUID type
		field.String("token").Unique(),
		field.String("ip_address").Optional().Nillable(),
		field.String("user_agent").Optional().Nillable(),
		field.Time("expires_at"),
		field.Time("last_login_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),

		field.String("spotify_access_token").Optional().Nillable(),
		field.String("spotify_refresh_token").Optional().Nillable(),
		field.Time("spotify_expiry").Optional().Nillable(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions").Unique(),
	}
}
