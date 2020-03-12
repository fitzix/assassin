package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Hot holds the schema definition for the Hot entity.
type Hot struct {
	ent.Schema
}

// Fields of the Hot.
func (Hot) Fields() []ent.Field {
	return []ent.Field{
		field.Int("hot").Default(0),
		field.Int("view").Default(0),
	}
}

// Edges of the Hot.
func (Hot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("hot").Unique().Required().StructTag(`json:"app,omitempty"`),
	}
}

func (Hot) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("app"),
	}
}
