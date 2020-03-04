package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty(),
		field.String("password").NotEmpty().Sensitive(),
		field.Uint("code").Positive().Default(1),
		field.Enum("status").Values("normal", "abnormal").Default("normal").Comment("normal: 正常; abnormal: 下架;"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).Unique().Required(),
	}
}