package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty().Unique(),
		field.String("icon").MaxLen(50).NotEmpty(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("apps", App.Type).Ref("categories"),
	}
}
