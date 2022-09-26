package schema

import (
	"context"
	"errors"
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gen "github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/ent/hook"

	"github.com/201R/go_api_boilerplate/packages/helpers"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable().DefaultFunc(uuid.NewString),
		field.String("firstName").
			Optional().
			Nillable().
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if len(s) > 15 {
					return errors.New("first name too long")
				}
				return nil
			}),
		field.String("lastName").Optional().Nillable().Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.String("email").NotEmpty().Unique(),
		field.String("location").Optional().Nillable(),
		field.String("phone").Optional().Nillable(),
		field.String("password").Sensitive().NotEmpty(),
		field.Bool("status").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					if password, ok := m.Password(); ok {
						hashedPassword, _ := helpers.Hash(password)
						m.SetPassword(string(hashedPassword))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate,
		),
	}
}
