package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Version holds the schema definition for the Version entity.
type Version struct {
	ent.Schema
}

func (Version) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeCreatedMixin{},
	}
}

// Fields of the Version.
func (Version) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("version"),
		field.String("size"),
		field.Int8("status").NonNegative().Default(1).Comment("0: 异常, 1: 正常"),
	}
}

// Edges of the Version.
func (Version) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("versions").Unique(),
		edge.To("sources", Source.Type),
	}
}

func (Version) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("version").Edges("app").Unique(),
	}
}
