package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

// Fields of the Announcement.
func (Announcement) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("description").
			NotEmpty(),
		field.String("time").
			NotEmpty(),
	}
}

// Edges of the Announcement.
func (Announcement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_announcement", Ukm.Type).
			Ref("announcement").
			Unique(),
	}
}
