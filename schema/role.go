package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty().Unique(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
