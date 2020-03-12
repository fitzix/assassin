// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/fitzix/assassin/ent/app"
	"github.com/fitzix/assassin/ent/hot"
)

// Hot is the model entity for the Hot schema.
type Hot struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hot holds the value of the "hot" field.
	Hot int `json:"hot,omitempty"`
	// View holds the value of the "view" field.
	View int `json:"view,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HotQuery when eager-loading is set.
	Edges   HotEdges `json:"edges"`
	app_hot *int
}

// HotEdges holds the relations/edges for other nodes in the graph.
type HotEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HotEdges) AppOrErr() (*App, error) {
	if e.loadedTypes[0] {
		if e.App == nil {
			// The edge app was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: app.Label}
		}
		return e.App, nil
	}
	return nil, &NotLoadedError{edge: "app"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hot) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // hot
		&sql.NullInt64{}, // view
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Hot) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // app_hot
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hot fields.
func (h *Hot) assignValues(values ...interface{}) error {
	if m, n := len(values), len(hot.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	h.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field hot", values[0])
	} else if value.Valid {
		h.Hot = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field view", values[1])
	} else if value.Valid {
		h.View = int(value.Int64)
	}
	values = values[2:]
	if len(values) == len(hot.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field app_hot", value)
		} else if value.Valid {
			h.app_hot = new(int)
			*h.app_hot = int(value.Int64)
		}
	}
	return nil
}

// QueryApp queries the app edge of the Hot.
func (h *Hot) QueryApp() *AppQuery {
	return (&HotClient{config: h.config}).QueryApp(h)
}

// Update returns a builder for updating this Hot.
// Note that, you need to call Hot.Unwrap() before calling this method, if this Hot
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hot) Update() *HotUpdateOne {
	return (&HotClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (h *Hot) Unwrap() *Hot {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hot is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hot) String() string {
	var builder strings.Builder
	builder.WriteString("Hot(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", hot=")
	builder.WriteString(fmt.Sprintf("%v", h.Hot))
	builder.WriteString(", view=")
	builder.WriteString(fmt.Sprintf("%v", h.View))
	builder.WriteByte(')')
	return builder.String()
}

// Hots is a parsable slice of Hot.
type Hots []*Hot

func (h Hots) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
