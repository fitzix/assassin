package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Vendor holds the schema definition for the Vendor entity.
type Provider struct {
	ent.Schema
}

// Fields of the Vendor.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Vendor.
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sources", Source.Type),
	}
}
