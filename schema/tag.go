// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schema

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Tag is an object representing the database table.
type Tag struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *tagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TagColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// Generated where

var TagWhere = struct {
	ID   whereHelperint
	Name whereHelperstring
}{
	ID:   whereHelperint{field: "\"tag\".\"id\""},
	Name: whereHelperstring{field: "\"tag\".\"name\""},
}

// TagRels is where relationship names are stored.
var TagRels = struct {
	Apps string
}{
	Apps: "Apps",
}

// tagR is where relationships are stored.
type tagR struct {
	Apps AppSlice
}

// NewStruct creates a new relationship struct
func (*tagR) NewStruct() *tagR {
	return &tagR{}
}

// tagL is where Load methods for each relationship are stored.
type tagL struct{}

var (
	tagAllColumns            = []string{"id", "name"}
	tagColumnsWithoutDefault = []string{"name"}
	tagColumnsWithDefault    = []string{"id"}
	tagPrimaryKeyColumns     = []string{"id"}
)

type (
	// TagSlice is an alias for a slice of pointers to Tag.
	// This should generally be used opposed to []Tag.
	TagSlice []*Tag

	tagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagType                 = reflect.TypeOf(&Tag{})
	tagMapping              = queries.MakeStructMapping(tagType)
	tagPrimaryKeyMapping, _ = queries.BindMapping(tagType, tagMapping, tagPrimaryKeyColumns)
	tagInsertCacheMut       sync.RWMutex
	tagInsertCache          = make(map[string]insertCache)
	tagUpdateCacheMut       sync.RWMutex
	tagUpdateCache          = make(map[string]updateCache)
	tagUpsertCacheMut       sync.RWMutex
	tagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single tag record from the query.
func (q tagQuery) One(exec boil.Executor) (*Tag, error) {
	o := &Tag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for tag")
	}

	return o, nil
}

// All returns all Tag records from the query.
func (q tagQuery) All(exec boil.Executor) (TagSlice, error) {
	var o []*Tag

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to Tag slice")
	}

	return o, nil
}

// Count returns the count of all Tag records in the query.
func (q tagQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count tag rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tagQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if tag exists")
	}

	return count > 0, nil
}

// Apps retrieves all the app's Apps with an executor.
func (o *Tag) Apps(mods ...qm.QueryMod) appQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"app_tag\" on \"app\".\"id\" = \"app_tag\".\"app_id\""),
		qm.Where("\"app_tag\".\"tag_id\"=?", o.ID),
	)

	query := Apps(queryMods...)
	queries.SetFrom(query.Query, "\"app\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"app\".*"})
	}

	return query
}

// LoadApps allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tagL) LoadApps(e boil.Executor, singular bool, maybeTag interface{}, mods queries.Applicator) error {
	var slice []*Tag
	var object *Tag

	if singular {
		object = maybeTag.(*Tag)
	} else {
		slice = *maybeTag.(*[]*Tag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tagR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"app\".*, \"a\".\"tag_id\""),
		qm.From("\"app\""),
		qm.InnerJoin("\"app_tag\" as \"a\" on \"app\".\"id\" = \"a\".\"app_id\""),
		qm.WhereIn("\"a\".\"tag_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load app")
	}

	var resultSlice []*App

	var localJoinCols []int
	for results.Next() {
		one := new(App)
		var localJoinCol int

		err = results.Scan(&one.ID, &one.HotID, &one.Name, &one.Type, &one.Icon, &one.Title, &one.Category, &one.CreatedAt, &one.UpdatedAt, &one.DeletedAt, &one.Status, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for app")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice app")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on app")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for app")
	}

	if singular {
		object.R.Apps = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &appR{}
			}
			foreign.R.Tags = append(foreign.R.Tags, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Apps = append(local.R.Apps, foreign)
				if foreign.R == nil {
					foreign.R = &appR{}
				}
				foreign.R.Tags = append(foreign.R.Tags, local)
				break
			}
		}
	}

	return nil
}

// AddApps adds the given related objects to the existing relationships
// of the tag, optionally inserting them as new records.
// Appends related to o.R.Apps.
// Sets related.R.Tags appropriately.
func (o *Tag) AddApps(exec boil.Executor, insert bool, related ...*App) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"app_tag\" (\"tag_id\", \"app_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &tagR{
			Apps: related,
		}
	} else {
		o.R.Apps = append(o.R.Apps, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &appR{
				Tags: TagSlice{o},
			}
		} else {
			rel.R.Tags = append(rel.R.Tags, o)
		}
	}
	return nil
}

// SetApps removes all previously related items of the
// tag replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Tags's Apps accordingly.
// Replaces o.R.Apps with related.
// Sets related.R.Tags's Apps accordingly.
func (o *Tag) SetApps(exec boil.Executor, insert bool, related ...*App) error {
	query := "delete from \"app_tag\" where \"tag_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeAppsFromTagsSlice(o, related)
	if o.R != nil {
		o.R.Apps = nil
	}
	return o.AddApps(exec, insert, related...)
}

// RemoveApps relationships from objects passed in.
// Removes related items from R.Apps (uses pointer comparison, removal does not keep order)
// Sets related.R.Tags.
func (o *Tag) RemoveApps(exec boil.Executor, related ...*App) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"app_tag\" where \"tag_id\" = $1 and \"app_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeAppsFromTagsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Apps {
			if rel != ri {
				continue
			}

			ln := len(o.R.Apps)
			if ln > 1 && i < ln-1 {
				o.R.Apps[i] = o.R.Apps[ln-1]
			}
			o.R.Apps = o.R.Apps[:ln-1]
			break
		}
	}

	return nil
}

func removeAppsFromTagsSlice(o *Tag, related []*App) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Tags {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Tags)
			if ln > 1 && i < ln-1 {
				rel.R.Tags[i] = rel.R.Tags[ln-1]
			}
			rel.R.Tags = rel.R.Tags[:ln-1]
			break
		}
	}
}

// Tags retrieves all the records using an executor.
func Tags(mods ...qm.QueryMod) tagQuery {
	mods = append(mods, qm.From("\"tag\""))
	return tagQuery{NewQuery(mods...)}
}

// FindTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTag(exec boil.Executor, iD int, selectCols ...string) (*Tag, error) {
	tagObj := &Tag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tag\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, tagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from tag")
	}

	return tagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tag) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no tag provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagInsertCacheMut.RLock()
	cache, cached := tagInsertCache[key]
	tagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagType, tagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tag\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tag\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "schema: unable to insert into tag")
	}

	if !cached {
		tagInsertCacheMut.Lock()
		tagInsertCache[key] = cache
		tagInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Tag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tag) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	tagUpdateCacheMut.RLock()
	cache, cached := tagUpdateCache[key]
	tagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update tag, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tag\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, append(wl, tagPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update tag row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for tag")
	}

	if !cached {
		tagUpdateCacheMut.Lock()
		tagUpdateCache[key] = cache
		tagUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q tagQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for tag")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("schema: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tag\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tagPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all tag")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tag) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no tag provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tagUpsertCacheMut.RLock()
	cache, cached := tagUpsertCache[key]
	tagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert tag, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tagPrimaryKeyColumns))
			copy(conflict, tagPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tag\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagType, tagMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "schema: unable to upsert tag")
	}

	if !cached {
		tagUpsertCacheMut.Lock()
		tagUpsertCache[key] = cache
		tagUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Tag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tag) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no Tag provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagPrimaryKeyMapping)
	sql := "DELETE FROM \"tag\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for tag")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tagQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no tagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for tag")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tag\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tagPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for tag")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tag) Reload(exec boil.Executor) error {
	ret, err := FindTag(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tag\".* FROM \"tag\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in TagSlice")
	}

	*o = slice

	return nil
}

// TagExists checks if the Tag row exists.
func TagExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tag\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if tag exists")
	}

	return exists, nil
}
