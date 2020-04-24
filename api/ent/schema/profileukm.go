package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ProfileUKM holds the schema definition for the ProfileUKM entity.
type ProfileUKM struct {
	ent.Schema
}

// Fields of the ProfileUKM.
func (ProfileUKM) Fields() []ent.Field {
	return []ent.Field{
		field.String("reason").
			NotEmpty(),
	}
}

// Config holds configuration of schema
func (ProfileUKM) Config() ent.Config {
	return ent.Config{
		Table: "profile_ukm",
	}
}

// Edges of the ProfileUKM.
func (ProfileUKM) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_profile", Profile.Type).
			Ref("ukms").
			Unique(),
		edge.From("owner_ukm", Ukm.Type).
			Ref("profiles").
			Unique(),
		edge.From("owner_role", RoleUKM.Type).
			Ref("profile_roles").
			Unique(),
	}
}
