// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/app"
	"github.com/fitzix/assassin/ent/source"
	"github.com/fitzix/assassin/ent/version"
)

// VersionCreate is the builder for creating a Version entity.
type VersionCreate struct {
	config
	created_at *time.Time
	name       *string
	version    *string
	size       *string
	status     *version.Status
	app        map[int]struct{}
	sources    map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (vc *VersionCreate) SetCreatedAt(t time.Time) *VersionCreate {
	vc.created_at = &t
	return vc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (vc *VersionCreate) SetNillableCreatedAt(t *time.Time) *VersionCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetName sets the name field.
func (vc *VersionCreate) SetName(s string) *VersionCreate {
	vc.name = &s
	return vc
}

// SetVersion sets the version field.
func (vc *VersionCreate) SetVersion(s string) *VersionCreate {
	vc.version = &s
	return vc
}

// SetSize sets the size field.
func (vc *VersionCreate) SetSize(s string) *VersionCreate {
	vc.size = &s
	return vc
}

// SetStatus sets the status field.
func (vc *VersionCreate) SetStatus(v version.Status) *VersionCreate {
	vc.status = &v
	return vc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (vc *VersionCreate) SetNillableStatus(v *version.Status) *VersionCreate {
	if v != nil {
		vc.SetStatus(*v)
	}
	return vc
}

// SetAppID sets the app edge to App by id.
func (vc *VersionCreate) SetAppID(id int) *VersionCreate {
	if vc.app == nil {
		vc.app = make(map[int]struct{})
	}
	vc.app[id] = struct{}{}
	return vc
}

// SetNillableAppID sets the app edge to App by id if the given value is not nil.
func (vc *VersionCreate) SetNillableAppID(id *int) *VersionCreate {
	if id != nil {
		vc = vc.SetAppID(*id)
	}
	return vc
}

// SetApp sets the app edge to App.
func (vc *VersionCreate) SetApp(a *App) *VersionCreate {
	return vc.SetAppID(a.ID)
}

// AddSourceIDs adds the sources edge to Source by ids.
func (vc *VersionCreate) AddSourceIDs(ids ...int) *VersionCreate {
	if vc.sources == nil {
		vc.sources = make(map[int]struct{})
	}
	for i := range ids {
		vc.sources[ids[i]] = struct{}{}
	}
	return vc
}

// AddSources adds the sources edges to Source.
func (vc *VersionCreate) AddSources(s ...*Source) *VersionCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vc.AddSourceIDs(ids...)
}

// Save creates the Version in the database.
func (vc *VersionCreate) Save(ctx context.Context) (*Version, error) {
	if vc.created_at == nil {
		v := version.DefaultCreatedAt()
		vc.created_at = &v
	}
	if vc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if vc.version == nil {
		return nil, errors.New("ent: missing required field \"version\"")
	}
	if vc.size == nil {
		return nil, errors.New("ent: missing required field \"size\"")
	}
	if vc.status == nil {
		v := version.DefaultStatus
		vc.status = &v
	}
	if err := version.StatusValidator(*vc.status); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
	}
	if len(vc.app) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"app\"")
	}
	return vc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VersionCreate) SaveX(ctx context.Context) *Version {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vc *VersionCreate) sqlSave(ctx context.Context) (*Version, error) {
	var (
		v     = &Version{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: version.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: version.FieldID,
			},
		}
	)
	if value := vc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: version.FieldCreatedAt,
		})
		v.CreatedAt = *value
	}
	if value := vc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: version.FieldName,
		})
		v.Name = *value
	}
	if value := vc.version; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: version.FieldVersion,
		})
		v.Version = *value
	}
	if value := vc.size; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: version.FieldSize,
		})
		v.Size = *value
	}
	if value := vc.status; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: version.FieldStatus,
		})
		v.Status = *value
	}
	if nodes := vc.app; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   version.AppTable,
			Columns: []string{version.AppColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.sources; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   version.SourcesTable,
			Columns: []string{version.SourcesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	v.ID = int(id)
	return v, nil
}
