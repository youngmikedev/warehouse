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
		field.String("access_token").NotEmpty(),
		field.String("refresh_token").NotEmpty(),
		// field.Int("expires_at_min"),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
		field.Bool("disabled").Default(false),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("sessions").
			Unique().
			Required(),
	}
}
