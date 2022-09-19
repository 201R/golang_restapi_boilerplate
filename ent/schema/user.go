package schema

import (
	"errors"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstName").
			Optional().
			Nillable().
			Unique().
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if len(s) > 15 {
					return errors.New("first name too long")
				}
				return nil
			}),
		field.String("lastName").Optional().Nillable().Unique().Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.String("email").NotEmpty().Unique(),
		field.String("location").Optional().Nillable(),
		field.String("phone").Optional().Nillable(),
		field.String("password").Sensitive().NotEmpty(),
		field.Bool("status").Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
