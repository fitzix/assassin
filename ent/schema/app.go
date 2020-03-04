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
		field.Enum("type").Values("app", "book").Default("app").Comment("app: 应用; book: 书籍;"),
		field.String("icon").NotEmpty(),
		field.String("title").MaxLen(200).Optional().Comment("标题"),
		field.Enum("status").Values("normal", "abnormal").Default("normal").Comment("normal: 正常; abnormal: 下架;"),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("categories", Category.Type),
		edge.To("carousels", Carousel.Type),
		edge.To("versions", Version.Type),
		edge.To("hot", Hot.Type).Unique(),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "type", "status", "deleted_at"),
	}
}
