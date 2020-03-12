// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/carousel"
	"github.com/fitzix/assassin/ent/predicate"
)

// CarouselDelete is the builder for deleting a Carousel entity.
type CarouselDelete struct {
	config
	predicates []predicate.Carousel
}

// Where adds a new predicate to the delete builder.
func (cd *CarouselDelete) Where(ps ...predicate.Carousel) *CarouselDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CarouselDelete) Exec(ctx context.Context) (int, error) {
	return cd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CarouselDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CarouselDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: carousel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carousel.FieldID,
			},
		},
	}
	if ps := cd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
}

// CarouselDeleteOne is the builder for deleting a single Carousel entity.
type CarouselDeleteOne struct {
	cd *CarouselDelete
}

// Exec executes the deletion query.
func (cdo *CarouselDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{carousel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CarouselDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
