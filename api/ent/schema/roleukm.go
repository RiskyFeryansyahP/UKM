package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// RoleUKM holds the schema definition for the RoleUKM entity.
type RoleUKM struct {
	ent.Schema
}

// Config holds configuration of schema
func (RoleUKM) Config() ent.Config {
	return ent.Config{
		Table: "role_ukm",
	}
}

// Fields of the RoleUKM.
func (RoleUKM) Fields() []ent.Field {
	return []ent.Field{
		field.String("status_role").
			NotEmpty(),
	}
}

// Edges of the RoleUKM.
func (RoleUKM) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile_roles", ProfileUKM.Type),
	}
}
