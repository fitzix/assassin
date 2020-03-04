// Code generated by entc, DO NOT EDIT.

package version

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/fitzix/assassin/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVersion), v))
	})
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVersion), v))
	})
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVersion), v))
	})
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVersion), v))
	})
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVersion), v))
	})
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...string) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...string) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// SizeContains applies the Contains predicate on the "size" field.
func SizeContains(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSize), v))
	})
}

// SizeHasPrefix applies the HasPrefix predicate on the "size" field.
func SizeHasPrefix(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSize), v))
	})
}

// SizeHasSuffix applies the HasSuffix predicate on the "size" field.
func SizeHasSuffix(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSize), v))
	})
}

// SizeEqualFold applies the EqualFold predicate on the "size" field.
func SizeEqualFold(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSize), v))
	})
}

// SizeContainsFold applies the ContainsFold predicate on the "size" field.
func SizeContainsFold(v string) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSize), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSources applies the HasEdge predicate on the "sources" edge.
func HasSources() predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SourcesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SourcesTable, SourcesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourcesWith applies the HasEdge predicate on the "sources" edge with a given conditions (other predicates).
func HasSourcesWith(preds ...predicate.Source) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SourcesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SourcesTable, SourcesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Version) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Version) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Version) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		p(s.Not())
	})
}