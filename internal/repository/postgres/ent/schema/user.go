package schema

import (
	"errors"
	"net/mail"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("password").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.String("email").Unique().NotEmpty().
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				if err != nil {
					return errors.New("invalid email address")
				}
				return nil
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("sessions", Session.Type),
	}
}
