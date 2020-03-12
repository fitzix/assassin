// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/predicate"
	"github.com/fitzix/assassin/ent/provider"
	"github.com/fitzix/assassin/ent/source"
)

// ProviderUpdate is the builder for updating Provider entities.
type ProviderUpdate struct {
	config
	name           *string
	sources        map[int]struct{}
	removedSources map[int]struct{}
	predicates     []predicate.Provider
}

// Where adds a new predicate for the builder.
func (pu *ProviderUpdate) Where(ps ...predicate.Provider) *ProviderUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *ProviderUpdate) SetName(s string) *ProviderUpdate {
	pu.name = &s
	return pu
}

// AddSourceIDs adds the sources edge to Source by ids.
func (pu *ProviderUpdate) AddSourceIDs(ids ...int) *ProviderUpdate {
	if pu.sources == nil {
		pu.sources = make(map[int]struct{})
	}
	for i := range ids {
		pu.sources[ids[i]] = struct{}{}
	}
	return pu
}

// AddSources adds the sources edges to Source.
func (pu *ProviderUpdate) AddSources(s ...*Source) *ProviderUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSourceIDs(ids...)
}

// RemoveSourceIDs removes the sources edge to Source by ids.
func (pu *ProviderUpdate) RemoveSourceIDs(ids ...int) *ProviderUpdate {
	if pu.removedSources == nil {
		pu.removedSources = make(map[int]struct{})
	}
	for i := range ids {
		pu.removedSources[ids[i]] = struct{}{}
	}
	return pu
}

// RemoveSources removes sources edges to Source.
func (pu *ProviderUpdate) RemoveSources(s ...*Source) *ProviderUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSourceIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProviderUpdate) Save(ctx context.Context) (int, error) {
	if pu.name != nil {
		if err := provider.NameValidator(*pu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProviderUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProviderUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   provider.Table,
			Columns: provider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: provider.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := pu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: provider.FieldName,
		})
	}
	if nodes := pu.removedSources; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.SourcesTable,
			Columns: []string{provider.SourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: source.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.sources; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.SourcesTable,
			Columns: []string{provider.SourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: source.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProviderUpdateOne is the builder for updating a single Provider entity.
type ProviderUpdateOne struct {
	config
	id             int
	name           *string
	sources        map[int]struct{}
	removedSources map[int]struct{}
}

// SetName sets the name field.
func (puo *ProviderUpdateOne) SetName(s string) *ProviderUpdateOne {
	puo.name = &s
	return puo
}

// AddSourceIDs adds the sources edge to Source by ids.
func (puo *ProviderUpdateOne) AddSourceIDs(ids ...int) *ProviderUpdateOne {
	if puo.sources == nil {
		puo.sources = make(map[int]struct{})
	}
	for i := range ids {
		puo.sources[ids[i]] = struct{}{}
	}
	return puo
}

// AddSources adds the sources edges to Source.
func (puo *ProviderUpdateOne) AddSources(s ...*Source) *ProviderUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSourceIDs(ids...)
}

// RemoveSourceIDs removes the sources edge to Source by ids.
func (puo *ProviderUpdateOne) RemoveSourceIDs(ids ...int) *ProviderUpdateOne {
	if puo.removedSources == nil {
		puo.removedSources = make(map[int]struct{})
	}
	for i := range ids {
		puo.removedSources[ids[i]] = struct{}{}
	}
	return puo
}

// RemoveSources removes sources edges to Source.
func (puo *ProviderUpdateOne) RemoveSources(s ...*Source) *ProviderUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSourceIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProviderUpdateOne) Save(ctx context.Context) (*Provider, error) {
	if puo.name != nil {
		if err := provider.NameValidator(*puo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProviderUpdateOne) SaveX(ctx context.Context) *Provider {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProviderUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProviderUpdateOne) sqlSave(ctx context.Context) (pr *Provider, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   provider.Table,
			Columns: provider.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  puo.id,
				Type:   field.TypeInt,
				Column: provider.FieldID,
			},
		},
	}
	if value := puo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: provider.FieldName,
		})
	}
	if nodes := puo.removedSources; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.SourcesTable,
			Columns: []string{provider.SourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: source.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.sources; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.SourcesTable,
			Columns: []string{provider.SourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: source.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Provider{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
