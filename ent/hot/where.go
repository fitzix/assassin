// Code generated by entc, DO NOT EDIT.

package hot

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/fitzix/assassin/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hot applies equality check predicate on the "hot" field. It's identical to HotEQ.
func Hot(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHot), v))
	})
}

// View applies equality check predicate on the "view" field. It's identical to ViewEQ.
func View(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldView), v))
	})
}

// HotEQ applies the EQ predicate on the "hot" field.
func HotEQ(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHot), v))
	})
}

// HotNEQ applies the NEQ predicate on the "hot" field.
func HotNEQ(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHot), v))
	})
}

// HotIn applies the In predicate on the "hot" field.
func HotIn(vs ...int) predicate.Hot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHot), v...))
	})
}

// HotNotIn applies the NotIn predicate on the "hot" field.
func HotNotIn(vs ...int) predicate.Hot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHot), v...))
	})
}

// HotGT applies the GT predicate on the "hot" field.
func HotGT(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHot), v))
	})
}

// HotGTE applies the GTE predicate on the "hot" field.
func HotGTE(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHot), v))
	})
}

// HotLT applies the LT predicate on the "hot" field.
func HotLT(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHot), v))
	})
}

// HotLTE applies the LTE predicate on the "hot" field.
func HotLTE(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHot), v))
	})
}

// ViewEQ applies the EQ predicate on the "view" field.
func ViewEQ(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldView), v))
	})
}

// ViewNEQ applies the NEQ predicate on the "view" field.
func ViewNEQ(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldView), v))
	})
}

// ViewIn applies the In predicate on the "view" field.
func ViewIn(vs ...int) predicate.Hot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldView), v...))
	})
}

// ViewNotIn applies the NotIn predicate on the "view" field.
func ViewNotIn(vs ...int) predicate.Hot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldView), v...))
	})
}

// ViewGT applies the GT predicate on the "view" field.
func ViewGT(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldView), v))
	})
}

// ViewGTE applies the GTE predicate on the "view" field.
func ViewGTE(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldView), v))
	})
}

// ViewLT applies the LT predicate on the "view" field.
func ViewLT(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldView), v))
	})
}

// ViewLTE applies the LTE predicate on the "view" field.
func ViewLTE(v int) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldView), v))
	})
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Hot) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Hot) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
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
func Not(p predicate.Hot) predicate.Hot {
	return predicate.Hot(func(s *sql.Selector) {
		p(s.Not())
	})
}
