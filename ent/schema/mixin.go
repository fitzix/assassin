package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type TimeCreatedMixin struct{}
type TimeUpdatedMixin struct{}
type TimeDeletedMixin struct{}
type TimeMixin struct{}

func (TimeCreatedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

func (TimeUpdatedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (TimeDeletedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").Optional(),
	}
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}
