// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/fitzix/assassin/ent/role"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(values ...interface{}) error {
	if m, n := len(values), len(role.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		r.Name = value.String
	}
	return nil
}

// Update returns a builder for updating this Role.
// Note that, you need to call Role.Unwrap() before calling this method, if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
