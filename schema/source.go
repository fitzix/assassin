package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

type Source struct {
	ent.Schema
}

// Fields of the Source.
func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").NotEmpty(),
		field.String("secret").Optional(),
	}
}

// Edges of the Source.
func (Source) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", Provider.Type).Ref("sources").Unique().Required(),
		edge.From("version", Version.Type).Ref("sources").Unique().Required(),
	}
}

func (Source) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("provider", "version"),
	}
}
