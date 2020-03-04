// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/predicate"
	"github.com/fitzix/assassin/ent/role"
	"github.com/fitzix/assassin/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	name        *string
	password    *string
	code        *uint
	addcode     *uint
	status      *user.Status
	role        map[int]struct{}
	clearedRole bool
	predicates  []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetPassword sets the password field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.password = &s
	return uu
}

// SetCode sets the code field.
func (uu *UserUpdate) SetCode(u uint) *UserUpdate {
	uu.code = &u
	uu.addcode = nil
	return uu
}

// SetNillableCode sets the code field if the given value is not nil.
func (uu *UserUpdate) SetNillableCode(u *uint) *UserUpdate {
	if u != nil {
		uu.SetCode(*u)
	}
	return uu
}

// AddCode adds u to code.
func (uu *UserUpdate) AddCode(u uint) *UserUpdate {
	if uu.addcode == nil {
		uu.addcode = &u
	} else {
		*uu.addcode += u
	}
	return uu
}

// SetStatus sets the status field.
func (uu *UserUpdate) SetStatus(u user.Status) *UserUpdate {
	uu.status = &u
	return uu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(u *user.Status) *UserUpdate {
	if u != nil {
		uu.SetStatus(*u)
	}
	return uu
}

// SetRoleID sets the role edge to Role by id.
func (uu *UserUpdate) SetRoleID(id int) *UserUpdate {
	if uu.role == nil {
		uu.role = make(map[int]struct{})
	}
	uu.role[id] = struct{}{}
	return uu
}

// SetRole sets the role edge to Role.
func (uu *UserUpdate) SetRole(r *Role) *UserUpdate {
	return uu.SetRoleID(r.ID)
}

// ClearRole clears the role edge to Role.
func (uu *UserUpdate) ClearRole() *UserUpdate {
	uu.clearedRole = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.name != nil {
		if err := user.NameValidator(*uu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if uu.password != nil {
		if err := user.PasswordValidator(*uu.password); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"password\": %v", err)
		}
	}
	if uu.code != nil {
		if err := user.CodeValidator(*uu.code); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"code\": %v", err)
		}
	}
	if uu.status != nil {
		if err := user.StatusValidator(*uu.status); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	if len(uu.role) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"role\"")
	}
	if uu.clearedRole && uu.role == nil {
		return 0, errors.New("ent: clearing a unique edge \"role\"")
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uu.password; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPassword,
		})
	}
	if value := uu.code; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: user.FieldCode,
		})
	}
	if value := uu.addcode; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: user.FieldCode,
		})
	}
	if value := uu.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldStatus,
		})
	}
	if uu.clearedRole {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.role; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id          int
	name        *string
	password    *string
	code        *uint
	addcode     *uint
	status      *user.Status
	role        map[int]struct{}
	clearedRole bool
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetPassword sets the password field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.password = &s
	return uuo
}

// SetCode sets the code field.
func (uuo *UserUpdateOne) SetCode(u uint) *UserUpdateOne {
	uuo.code = &u
	uuo.addcode = nil
	return uuo
}

// SetNillableCode sets the code field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCode(u *uint) *UserUpdateOne {
	if u != nil {
		uuo.SetCode(*u)
	}
	return uuo
}

// AddCode adds u to code.
func (uuo *UserUpdateOne) AddCode(u uint) *UserUpdateOne {
	if uuo.addcode == nil {
		uuo.addcode = &u
	} else {
		*uuo.addcode += u
	}
	return uuo
}

// SetStatus sets the status field.
func (uuo *UserUpdateOne) SetStatus(u user.Status) *UserUpdateOne {
	uuo.status = &u
	return uuo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(u *user.Status) *UserUpdateOne {
	if u != nil {
		uuo.SetStatus(*u)
	}
	return uuo
}

// SetRoleID sets the role edge to Role by id.
func (uuo *UserUpdateOne) SetRoleID(id int) *UserUpdateOne {
	if uuo.role == nil {
		uuo.role = make(map[int]struct{})
	}
	uuo.role[id] = struct{}{}
	return uuo
}

// SetRole sets the role edge to Role.
func (uuo *UserUpdateOne) SetRole(r *Role) *UserUpdateOne {
	return uuo.SetRoleID(r.ID)
}

// ClearRole clears the role edge to Role.
func (uuo *UserUpdateOne) ClearRole() *UserUpdateOne {
	uuo.clearedRole = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.name != nil {
		if err := user.NameValidator(*uuo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if uuo.password != nil {
		if err := user.PasswordValidator(*uuo.password); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"password\": %v", err)
		}
	}
	if uuo.code != nil {
		if err := user.CodeValidator(*uuo.code); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"code\": %v", err)
		}
	}
	if uuo.status != nil {
		if err := user.StatusValidator(*uuo.status); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	if len(uuo.role) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"role\"")
	}
	if uuo.clearedRole && uuo.role == nil {
		return nil, errors.New("ent: clearing a unique edge \"role\"")
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uuo.password; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPassword,
		})
	}
	if value := uuo.code; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: user.FieldCode,
		})
	}
	if value := uuo.addcode; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: user.FieldCode,
		})
	}
	if value := uuo.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldStatus,
		})
	}
	if uuo.clearedRole {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.role; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
