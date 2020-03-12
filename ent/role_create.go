// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/role"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	name *string
}

// SetName sets the name field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.name = &s
	return rc
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	if rc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := role.NameValidator(*rc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	return rc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	var (
		r     = &Role{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: role.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		}
	)
	if value := rc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: role.FieldName,
		})
		r.Name = *value
	}
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}
