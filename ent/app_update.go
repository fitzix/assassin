// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/fitzix/assassin/ent/app"
	"github.com/fitzix/assassin/ent/carousel"
	"github.com/fitzix/assassin/ent/category"
	"github.com/fitzix/assassin/ent/hot"
	"github.com/fitzix/assassin/ent/predicate"
	"github.com/fitzix/assassin/ent/tag"
	"github.com/fitzix/assassin/ent/version"
)

// AppUpdate is the builder for updating App entities.
type AppUpdate struct {
	config

	updated_at        *time.Time
	deleted_at        *time.Time
	cleardeleted_at   bool
	name              *string
	_type             *app.Type
	icon              *string
	title             *string
	cleartitle        bool
	status            *app.Status
	tags              map[int]struct{}
	categories        map[int]struct{}
	carousels         map[int]struct{}
	versions          map[int]struct{}
	hot               map[int]struct{}
	removedTags       map[int]struct{}
	removedCategories map[int]struct{}
	removedCarousels  map[int]struct{}
	removedVersions   map[int]struct{}
	clearedHot        bool
	predicates        []predicate.App
}

// Where adds a new predicate for the builder.
func (au *AppUpdate) Where(ps ...predicate.App) *AppUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetUpdatedAt sets the updated_at field.
func (au *AppUpdate) SetUpdatedAt(t time.Time) *AppUpdate {
	au.updated_at = &t
	return au
}

// SetDeletedAt sets the deleted_at field.
func (au *AppUpdate) SetDeletedAt(t time.Time) *AppUpdate {
	au.deleted_at = &t
	return au
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (au *AppUpdate) SetNillableDeletedAt(t *time.Time) *AppUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of deleted_at.
func (au *AppUpdate) ClearDeletedAt() *AppUpdate {
	au.deleted_at = nil
	au.cleardeleted_at = true
	return au
}

// SetName sets the name field.
func (au *AppUpdate) SetName(s string) *AppUpdate {
	au.name = &s
	return au
}

// SetType sets the type field.
func (au *AppUpdate) SetType(a app.Type) *AppUpdate {
	au._type = &a
	return au
}

// SetNillableType sets the type field if the given value is not nil.
func (au *AppUpdate) SetNillableType(a *app.Type) *AppUpdate {
	if a != nil {
		au.SetType(*a)
	}
	return au
}

// SetIcon sets the icon field.
func (au *AppUpdate) SetIcon(s string) *AppUpdate {
	au.icon = &s
	return au
}

// SetTitle sets the title field.
func (au *AppUpdate) SetTitle(s string) *AppUpdate {
	au.title = &s
	return au
}

// SetNillableTitle sets the title field if the given value is not nil.
func (au *AppUpdate) SetNillableTitle(s *string) *AppUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// ClearTitle clears the value of title.
func (au *AppUpdate) ClearTitle() *AppUpdate {
	au.title = nil
	au.cleartitle = true
	return au
}

// SetStatus sets the status field.
func (au *AppUpdate) SetStatus(a app.Status) *AppUpdate {
	au.status = &a
	return au
}

// SetNillableStatus sets the status field if the given value is not nil.
func (au *AppUpdate) SetNillableStatus(a *app.Status) *AppUpdate {
	if a != nil {
		au.SetStatus(*a)
	}
	return au
}

// AddTagIDs adds the tags edge to Tag by ids.
func (au *AppUpdate) AddTagIDs(ids ...int) *AppUpdate {
	if au.tags == nil {
		au.tags = make(map[int]struct{})
	}
	for i := range ids {
		au.tags[ids[i]] = struct{}{}
	}
	return au
}

// AddTags adds the tags edges to Tag.
func (au *AppUpdate) AddTags(t ...*Tag) *AppUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTagIDs(ids...)
}

// AddCategoryIDs adds the categories edge to Category by ids.
func (au *AppUpdate) AddCategoryIDs(ids ...int) *AppUpdate {
	if au.categories == nil {
		au.categories = make(map[int]struct{})
	}
	for i := range ids {
		au.categories[ids[i]] = struct{}{}
	}
	return au
}

// AddCategories adds the categories edges to Category.
func (au *AppUpdate) AddCategories(c ...*Category) *AppUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCategoryIDs(ids...)
}

// AddCarouselIDs adds the carousels edge to Carousel by ids.
func (au *AppUpdate) AddCarouselIDs(ids ...int) *AppUpdate {
	if au.carousels == nil {
		au.carousels = make(map[int]struct{})
	}
	for i := range ids {
		au.carousels[ids[i]] = struct{}{}
	}
	return au
}

// AddCarousels adds the carousels edges to Carousel.
func (au *AppUpdate) AddCarousels(c ...*Carousel) *AppUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCarouselIDs(ids...)
}

// AddVersionIDs adds the versions edge to Version by ids.
func (au *AppUpdate) AddVersionIDs(ids ...int) *AppUpdate {
	if au.versions == nil {
		au.versions = make(map[int]struct{})
	}
	for i := range ids {
		au.versions[ids[i]] = struct{}{}
	}
	return au
}

// AddVersions adds the versions edges to Version.
func (au *AppUpdate) AddVersions(v ...*Version) *AppUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.AddVersionIDs(ids...)
}

// SetHotID sets the hot edge to Hot by id.
func (au *AppUpdate) SetHotID(id int) *AppUpdate {
	if au.hot == nil {
		au.hot = make(map[int]struct{})
	}
	au.hot[id] = struct{}{}
	return au
}

// SetNillableHotID sets the hot edge to Hot by id if the given value is not nil.
func (au *AppUpdate) SetNillableHotID(id *int) *AppUpdate {
	if id != nil {
		au = au.SetHotID(*id)
	}
	return au
}

// SetHot sets the hot edge to Hot.
func (au *AppUpdate) SetHot(h *Hot) *AppUpdate {
	return au.SetHotID(h.ID)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (au *AppUpdate) RemoveTagIDs(ids ...int) *AppUpdate {
	if au.removedTags == nil {
		au.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		au.removedTags[ids[i]] = struct{}{}
	}
	return au
}

// RemoveTags removes tags edges to Tag.
func (au *AppUpdate) RemoveTags(t ...*Tag) *AppUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTagIDs(ids...)
}

// RemoveCategoryIDs removes the categories edge to Category by ids.
func (au *AppUpdate) RemoveCategoryIDs(ids ...int) *AppUpdate {
	if au.removedCategories == nil {
		au.removedCategories = make(map[int]struct{})
	}
	for i := range ids {
		au.removedCategories[ids[i]] = struct{}{}
	}
	return au
}

// RemoveCategories removes categories edges to Category.
func (au *AppUpdate) RemoveCategories(c ...*Category) *AppUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCategoryIDs(ids...)
}

// RemoveCarouselIDs removes the carousels edge to Carousel by ids.
func (au *AppUpdate) RemoveCarouselIDs(ids ...int) *AppUpdate {
	if au.removedCarousels == nil {
		au.removedCarousels = make(map[int]struct{})
	}
	for i := range ids {
		au.removedCarousels[ids[i]] = struct{}{}
	}
	return au
}

// RemoveCarousels removes carousels edges to Carousel.
func (au *AppUpdate) RemoveCarousels(c ...*Carousel) *AppUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCarouselIDs(ids...)
}

// RemoveVersionIDs removes the versions edge to Version by ids.
func (au *AppUpdate) RemoveVersionIDs(ids ...int) *AppUpdate {
	if au.removedVersions == nil {
		au.removedVersions = make(map[int]struct{})
	}
	for i := range ids {
		au.removedVersions[ids[i]] = struct{}{}
	}
	return au
}

// RemoveVersions removes versions edges to Version.
func (au *AppUpdate) RemoveVersions(v ...*Version) *AppUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.RemoveVersionIDs(ids...)
}

// ClearHot clears the hot edge to Hot.
func (au *AppUpdate) ClearHot() *AppUpdate {
	au.clearedHot = true
	return au
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AppUpdate) Save(ctx context.Context) (int, error) {
	if au.updated_at == nil {
		v := app.UpdateDefaultUpdatedAt()
		au.updated_at = &v
	}
	if au.name != nil {
		if err := app.NameValidator(*au.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if au._type != nil {
		if err := app.TypeValidator(*au._type); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	if au.icon != nil {
		if err := app.IconValidator(*au.icon); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"icon\": %v", err)
		}
	}
	if au.title != nil {
		if err := app.TitleValidator(*au.title); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}
	if au.status != nil {
		if err := app.StatusValidator(*au.status); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	if len(au.hot) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"hot\"")
	}
	return au.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := au.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: app.FieldUpdatedAt,
		})
	}
	if value := au.deleted_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: app.FieldDeletedAt,
		})
	}
	if au.cleardeleted_at {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: app.FieldDeletedAt,
		})
	}
	if value := au.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: app.FieldName,
		})
	}
	if value := au._type; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: app.FieldType,
		})
	}
	if value := au.icon; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: app.FieldIcon,
		})
	}
	if value := au.title; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: app.FieldTitle,
		})
	}
	if au.cleartitle {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: app.FieldTitle,
		})
	}
	if value := au.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: app.FieldStatus,
		})
	}
	if nodes := au.removedTags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.TagsTable,
			Columns: app.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.TagsTable,
			Columns: app.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := au.removedCategories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.CategoriesTable,
			Columns: app.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.categories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.CategoriesTable,
			Columns: app.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := au.removedCarousels; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.CarouselsTable,
			Columns: []string{app.CarouselsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carousel.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.carousels; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.CarouselsTable,
			Columns: []string{app.CarouselsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carousel.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := au.removedVersions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.VersionsTable,
			Columns: []string{app.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: version.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.versions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.VersionsTable,
			Columns: []string{app.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: version.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.clearedHot {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   app.HotTable,
			Columns: []string{app.HotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.hot; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   app.HotTable,
			Columns: []string{app.HotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hot.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AppUpdateOne is the builder for updating a single App entity.
type AppUpdateOne struct {
	config
	id int

	updated_at        *time.Time
	deleted_at        *time.Time
	cleardeleted_at   bool
	name              *string
	_type             *app.Type
	icon              *string
	title             *string
	cleartitle        bool
	status            *app.Status
	tags              map[int]struct{}
	categories        map[int]struct{}
	carousels         map[int]struct{}
	versions          map[int]struct{}
	hot               map[int]struct{}
	removedTags       map[int]struct{}
	removedCategories map[int]struct{}
	removedCarousels  map[int]struct{}
	removedVersions   map[int]struct{}
	clearedHot        bool
}

// SetUpdatedAt sets the updated_at field.
func (auo *AppUpdateOne) SetUpdatedAt(t time.Time) *AppUpdateOne {
	auo.updated_at = &t
	return auo
}

// SetDeletedAt sets the deleted_at field.
func (auo *AppUpdateOne) SetDeletedAt(t time.Time) *AppUpdateOne {
	auo.deleted_at = &t
	return auo
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableDeletedAt(t *time.Time) *AppUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of deleted_at.
func (auo *AppUpdateOne) ClearDeletedAt() *AppUpdateOne {
	auo.deleted_at = nil
	auo.cleardeleted_at = true
	return auo
}

// SetName sets the name field.
func (auo *AppUpdateOne) SetName(s string) *AppUpdateOne {
	auo.name = &s
	return auo
}

// SetType sets the type field.
func (auo *AppUpdateOne) SetType(a app.Type) *AppUpdateOne {
	auo._type = &a
	return auo
}

// SetNillableType sets the type field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableType(a *app.Type) *AppUpdateOne {
	if a != nil {
		auo.SetType(*a)
	}
	return auo
}

// SetIcon sets the icon field.
func (auo *AppUpdateOne) SetIcon(s string) *AppUpdateOne {
	auo.icon = &s
	return auo
}

// SetTitle sets the title field.
func (auo *AppUpdateOne) SetTitle(s string) *AppUpdateOne {
	auo.title = &s
	return auo
}

// SetNillableTitle sets the title field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableTitle(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// ClearTitle clears the value of title.
func (auo *AppUpdateOne) ClearTitle() *AppUpdateOne {
	auo.title = nil
	auo.cleartitle = true
	return auo
}

// SetStatus sets the status field.
func (auo *AppUpdateOne) SetStatus(a app.Status) *AppUpdateOne {
	auo.status = &a
	return auo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableStatus(a *app.Status) *AppUpdateOne {
	if a != nil {
		auo.SetStatus(*a)
	}
	return auo
}

// AddTagIDs adds the tags edge to Tag by ids.
func (auo *AppUpdateOne) AddTagIDs(ids ...int) *AppUpdateOne {
	if auo.tags == nil {
		auo.tags = make(map[int]struct{})
	}
	for i := range ids {
		auo.tags[ids[i]] = struct{}{}
	}
	return auo
}

// AddTags adds the tags edges to Tag.
func (auo *AppUpdateOne) AddTags(t ...*Tag) *AppUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTagIDs(ids...)
}

// AddCategoryIDs adds the categories edge to Category by ids.
func (auo *AppUpdateOne) AddCategoryIDs(ids ...int) *AppUpdateOne {
	if auo.categories == nil {
		auo.categories = make(map[int]struct{})
	}
	for i := range ids {
		auo.categories[ids[i]] = struct{}{}
	}
	return auo
}

// AddCategories adds the categories edges to Category.
func (auo *AppUpdateOne) AddCategories(c ...*Category) *AppUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCategoryIDs(ids...)
}

// AddCarouselIDs adds the carousels edge to Carousel by ids.
func (auo *AppUpdateOne) AddCarouselIDs(ids ...int) *AppUpdateOne {
	if auo.carousels == nil {
		auo.carousels = make(map[int]struct{})
	}
	for i := range ids {
		auo.carousels[ids[i]] = struct{}{}
	}
	return auo
}

// AddCarousels adds the carousels edges to Carousel.
func (auo *AppUpdateOne) AddCarousels(c ...*Carousel) *AppUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCarouselIDs(ids...)
}

// AddVersionIDs adds the versions edge to Version by ids.
func (auo *AppUpdateOne) AddVersionIDs(ids ...int) *AppUpdateOne {
	if auo.versions == nil {
		auo.versions = make(map[int]struct{})
	}
	for i := range ids {
		auo.versions[ids[i]] = struct{}{}
	}
	return auo
}

// AddVersions adds the versions edges to Version.
func (auo *AppUpdateOne) AddVersions(v ...*Version) *AppUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.AddVersionIDs(ids...)
}

// SetHotID sets the hot edge to Hot by id.
func (auo *AppUpdateOne) SetHotID(id int) *AppUpdateOne {
	if auo.hot == nil {
		auo.hot = make(map[int]struct{})
	}
	auo.hot[id] = struct{}{}
	return auo
}

// SetNillableHotID sets the hot edge to Hot by id if the given value is not nil.
func (auo *AppUpdateOne) SetNillableHotID(id *int) *AppUpdateOne {
	if id != nil {
		auo = auo.SetHotID(*id)
	}
	return auo
}

// SetHot sets the hot edge to Hot.
func (auo *AppUpdateOne) SetHot(h *Hot) *AppUpdateOne {
	return auo.SetHotID(h.ID)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (auo *AppUpdateOne) RemoveTagIDs(ids ...int) *AppUpdateOne {
	if auo.removedTags == nil {
		auo.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		auo.removedTags[ids[i]] = struct{}{}
	}
	return auo
}

// RemoveTags removes tags edges to Tag.
func (auo *AppUpdateOne) RemoveTags(t ...*Tag) *AppUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTagIDs(ids...)
}

// RemoveCategoryIDs removes the categories edge to Category by ids.
func (auo *AppUpdateOne) RemoveCategoryIDs(ids ...int) *AppUpdateOne {
	if auo.removedCategories == nil {
		auo.removedCategories = make(map[int]struct{})
	}
	for i := range ids {
		auo.removedCategories[ids[i]] = struct{}{}
	}
	return auo
}

// RemoveCategories removes categories edges to Category.
func (auo *AppUpdateOne) RemoveCategories(c ...*Category) *AppUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCategoryIDs(ids...)
}

// RemoveCarouselIDs removes the carousels edge to Carousel by ids.
func (auo *AppUpdateOne) RemoveCarouselIDs(ids ...int) *AppUpdateOne {
	if auo.removedCarousels == nil {
		auo.removedCarousels = make(map[int]struct{})
	}
	for i := range ids {
		auo.removedCarousels[ids[i]] = struct{}{}
	}
	return auo
}

// RemoveCarousels removes carousels edges to Carousel.
func (auo *AppUpdateOne) RemoveCarousels(c ...*Carousel) *AppUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCarouselIDs(ids...)
}

// RemoveVersionIDs removes the versions edge to Version by ids.
func (auo *AppUpdateOne) RemoveVersionIDs(ids ...int) *AppUpdateOne {
	if auo.removedVersions == nil {
		auo.removedVersions = make(map[int]struct{})
	}
	for i := range ids {
		auo.removedVersions[ids[i]] = struct{}{}
	}
	return auo
}

// RemoveVersions removes versions edges to Version.
func (auo *AppUpdateOne) RemoveVersions(v ...*Version) *AppUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.RemoveVersionIDs(ids...)
}

// ClearHot clears the hot edge to Hot.
func (auo *AppUpdateOne) ClearHot() *AppUpdateOne {
	auo.clearedHot = true
	return auo
}

// Save executes the query and returns the updated entity.
func (auo *AppUpdateOne) Save(ctx context.Context) (*App, error) {
	if auo.updated_at == nil {
		v := app.UpdateDefaultUpdatedAt()
		auo.updated_at = &v
	}
	if auo.name != nil {
		if err := app.NameValidator(*auo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if auo._type != nil {
		if err := app.TypeValidator(*auo._type); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	if auo.icon != nil {
		if err := app.IconValidator(*auo.icon); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"icon\": %v", err)
		}
	}
	if auo.title != nil {
		if err := app.TitleValidator(*auo.title); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
		}
	}
	if auo.status != nil {
		if err := app.StatusValidator(*auo.status); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	if len(auo.hot) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"hot\"")
	}
	return auo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppUpdateOne) SaveX(ctx context.Context) *App {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AppUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AppUpdateOne) sqlSave(ctx context.Context) (a *App, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  auo.id,
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
	}
	if value := auo.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: app.FieldUpdatedAt,
		})
	}
	if value := auo.deleted_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: app.FieldDeletedAt,
		})
	}
	if auo.cleardeleted_at {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: app.FieldDeletedAt,
		})
	}
	if value := auo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: app.FieldName,
		})
	}
	if value := auo._type; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: app.FieldType,
		})
	}
	if value := auo.icon; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: app.FieldIcon,
		})
	}
	if value := auo.title; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: app.FieldTitle,
		})
	}
	if auo.cleartitle {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: app.FieldTitle,
		})
	}
	if value := auo.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: app.FieldStatus,
		})
	}
	if nodes := auo.removedTags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.TagsTable,
			Columns: app.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.TagsTable,
			Columns: app.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := auo.removedCategories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.CategoriesTable,
			Columns: app.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.categories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   app.CategoriesTable,
			Columns: app.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := auo.removedCarousels; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.CarouselsTable,
			Columns: []string{app.CarouselsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carousel.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.carousels; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.CarouselsTable,
			Columns: []string{app.CarouselsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carousel.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := auo.removedVersions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.VersionsTable,
			Columns: []string{app.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: version.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.versions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.VersionsTable,
			Columns: []string{app.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: version.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.clearedHot {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   app.HotTable,
			Columns: []string{app.HotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.hot; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   app.HotTable,
			Columns: []string{app.HotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hot.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &App{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}