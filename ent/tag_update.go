// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/app"
	"github.com/fitzix/assassin/ent/predicate"
	"github.com/fitzix/assassin/ent/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	name       *string
	app        map[int]struct{}
	removedApp map[int]struct{}
	predicates []predicate.Tag
}

// Where adds a new predicate for the builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetName sets the name field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.name = &s
	return tu
}

// AddAppIDs adds the app edge to App by ids.
func (tu *TagUpdate) AddAppIDs(ids ...int) *TagUpdate {
	if tu.app == nil {
		tu.app = make(map[int]struct{})
	}
	for i := range ids {
		tu.app[ids[i]] = struct{}{}
	}
	return tu
}

// AddApp adds the app edges to App.
func (tu *TagUpdate) AddApp(a ...*App) *TagUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.AddAppIDs(ids...)
}

// RemoveAppIDs removes the app edge to App by ids.
func (tu *TagUpdate) RemoveAppIDs(ids ...int) *TagUpdate {
	if tu.removedApp == nil {
		tu.removedApp = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedApp[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveApp removes app edges to App.
func (tu *TagUpdate) RemoveApp(a ...*App) *TagUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.RemoveAppIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	if tu.name != nil {
		if err := tag.NameValidator(*tu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := tu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: tag.FieldName,
		})
	}
	if nodes := tu.removedApp; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.AppTable,
			Columns: tag.AppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.app; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.AppTable,
			Columns: tag.AppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	id         int
	name       *string
	app        map[int]struct{}
	removedApp map[int]struct{}
}

// SetName sets the name field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.name = &s
	return tuo
}

// AddAppIDs adds the app edge to App by ids.
func (tuo *TagUpdateOne) AddAppIDs(ids ...int) *TagUpdateOne {
	if tuo.app == nil {
		tuo.app = make(map[int]struct{})
	}
	for i := range ids {
		tuo.app[ids[i]] = struct{}{}
	}
	return tuo
}

// AddApp adds the app edges to App.
func (tuo *TagUpdateOne) AddApp(a ...*App) *TagUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.AddAppIDs(ids...)
}

// RemoveAppIDs removes the app edge to App by ids.
func (tuo *TagUpdateOne) RemoveAppIDs(ids ...int) *TagUpdateOne {
	if tuo.removedApp == nil {
		tuo.removedApp = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedApp[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveApp removes app edges to App.
func (tuo *TagUpdateOne) RemoveApp(a ...*App) *TagUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.RemoveAppIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	if tuo.name != nil {
		if err := tag.NameValidator(*tuo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (t *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  tuo.id,
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if value := tuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: tag.FieldName,
		})
	}
	if nodes := tuo.removedApp; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.AppTable,
			Columns: tag.AppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.app; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.AppTable,
			Columns: tag.AppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Tag{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
