package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Carousel holds the schema definition for the Carousel entity.
type Carousel struct {
	ent.Schema
}

// Fields of the Carousel.
func (Carousel) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").NotEmpty(),
	}
}

// Edges of the Carousel.
func (Carousel) Edges() []ent.Edge {
	return nil
}
