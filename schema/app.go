package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty().Unique(),
		field.Int8("type").NonNegative().Default(0).Comment("0: app, 1: book"),
		field.String("icon").NotEmpty(),
		field.String("title").MaxLen(200).Optional().Comment("标题"),
		field.Int8("status").NonNegative().Default(1).Comment("0: 异常, 1: 正常"),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).StructTag(`json:"tags,omitempty"`),
		edge.To("categories", Category.Type).StructTag(`json:"categories,omitempty"`),
		edge.To("carousels", Carousel.Type).StructTag(`json:"carousels,omitempty"`),
		edge.To("versions", Version.Type).StructTag(`json:"versions,omitempty"`),
		edge.To("hot", Hot.Type).Unique().StructTag(`json:"hot,omitempty"`),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "type", "status", "deleted_at"),
	}
}
