package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Ukm holds the schema definition for the Ukm entity.
type Ukm struct {
	ent.Schema
}

// Fields of the Ukm.
func (Ukm) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Enum("status").
			Values("open", "close"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Ukm.
func (Ukm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profiles", ProfileUKM.Type),
		edge.To("announcement", Announcement.Type),
	}
}
