package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstName").
			NotEmpty(),
		field.String("lastName").
			NotEmpty(),
		field.Enum("jk").
			Values("Male", "Female", "None"),
		field.String("address").
			Optional(),
		field.String("birth_date").
			Optional(),
		field.String("year_generation"),
		field.String("phone").
			NotEmpty(),
		field.Bool("status").
			Default(true),
		field.String("img_profile").
			Default(""),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("profile").
			Required(),
		edge.To("ukms", ProfileUKM.Type),
	}
}
