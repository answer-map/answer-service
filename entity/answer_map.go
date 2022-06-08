// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AnswerMap is an object representing the database table.
type AnswerMap struct {
	MapID       string      `boil:"map_id" json:"map_id" toml:"map_id" yaml:"map_id"`
	UserID      string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	AnswerKey   string      `boil:"answer_key" json:"answer_key" toml:"answer_key" yaml:"answer_key"`
	AnswerValue null.String `boil:"answer_value" json:"answer_value,omitempty" toml:"answer_value" yaml:"answer_value,omitempty"`

	R *answerMapR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L answerMapL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnswerMapColumns = struct {
	MapID       string
	UserID      string
	AnswerKey   string
	AnswerValue string
}{
	MapID:       "map_id",
	UserID:      "user_id",
	AnswerKey:   "answer_key",
	AnswerValue: "answer_value",
}

var AnswerMapTableColumns = struct {
	MapID       string
	UserID      string
	AnswerKey   string
	AnswerValue string
}{
	MapID:       "answer_map.map_id",
	UserID:      "answer_map.user_id",
	AnswerKey:   "answer_map.answer_key",
	AnswerValue: "answer_map.answer_value",
}

// Generated where

var AnswerMapWhere = struct {
	MapID       whereHelperstring
	UserID      whereHelperstring
	AnswerKey   whereHelperstring
	AnswerValue whereHelpernull_String
}{
	MapID:       whereHelperstring{field: "\"answer_map\".\"map_id\""},
	UserID:      whereHelperstring{field: "\"answer_map\".\"user_id\""},
	AnswerKey:   whereHelperstring{field: "\"answer_map\".\"answer_key\""},
	AnswerValue: whereHelpernull_String{field: "\"answer_map\".\"answer_value\""},
}

// AnswerMapRels is where relationship names are stored.
var AnswerMapRels = struct {
	User            string
	MapAnswerEvents string
}{
	User:            "User",
	MapAnswerEvents: "MapAnswerEvents",
}

// answerMapR is where relationships are stored.
type answerMapR struct {
	User            *AnswerUser      `boil:"User" json:"User" toml:"User" yaml:"User"`
	MapAnswerEvents AnswerEventSlice `boil:"MapAnswerEvents" json:"MapAnswerEvents" toml:"MapAnswerEvents" yaml:"MapAnswerEvents"`
}

// NewStruct creates a new relationship struct
func (*answerMapR) NewStruct() *answerMapR {
	return &answerMapR{}
}

// answerMapL is where Load methods for each relationship are stored.
type answerMapL struct{}

var (
	answerMapAllColumns            = []string{"map_id", "user_id", "answer_key", "answer_value"}
	answerMapColumnsWithoutDefault = []string{"map_id", "user_id", "answer_key"}
	answerMapColumnsWithDefault    = []string{"answer_value"}
	answerMapPrimaryKeyColumns     = []string{"map_id"}
	answerMapGeneratedColumns      = []string{}
)

type (
	// AnswerMapSlice is an alias for a slice of pointers to AnswerMap.
	// This should almost always be used instead of []AnswerMap.
	AnswerMapSlice []*AnswerMap
	// AnswerMapHook is the signature for custom AnswerMap hook methods
	AnswerMapHook func(context.Context, boil.ContextExecutor, *AnswerMap) error

	answerMapQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	answerMapType                 = reflect.TypeOf(&AnswerMap{})
	answerMapMapping              = queries.MakeStructMapping(answerMapType)
	answerMapPrimaryKeyMapping, _ = queries.BindMapping(answerMapType, answerMapMapping, answerMapPrimaryKeyColumns)
	answerMapInsertCacheMut       sync.RWMutex
	answerMapInsertCache          = make(map[string]insertCache)
	answerMapUpdateCacheMut       sync.RWMutex
	answerMapUpdateCache          = make(map[string]updateCache)
	answerMapUpsertCacheMut       sync.RWMutex
	answerMapUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var answerMapAfterSelectHooks []AnswerMapHook

var answerMapBeforeInsertHooks []AnswerMapHook
var answerMapAfterInsertHooks []AnswerMapHook

var answerMapBeforeUpdateHooks []AnswerMapHook
var answerMapAfterUpdateHooks []AnswerMapHook

var answerMapBeforeDeleteHooks []AnswerMapHook
var answerMapAfterDeleteHooks []AnswerMapHook

var answerMapBeforeUpsertHooks []AnswerMapHook
var answerMapAfterUpsertHooks []AnswerMapHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AnswerMap) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AnswerMap) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AnswerMap) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AnswerMap) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AnswerMap) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AnswerMap) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AnswerMap) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AnswerMap) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AnswerMap) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerMapAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnswerMapHook registers your hook function for all future operations.
func AddAnswerMapHook(hookPoint boil.HookPoint, answerMapHook AnswerMapHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		answerMapAfterSelectHooks = append(answerMapAfterSelectHooks, answerMapHook)
	case boil.BeforeInsertHook:
		answerMapBeforeInsertHooks = append(answerMapBeforeInsertHooks, answerMapHook)
	case boil.AfterInsertHook:
		answerMapAfterInsertHooks = append(answerMapAfterInsertHooks, answerMapHook)
	case boil.BeforeUpdateHook:
		answerMapBeforeUpdateHooks = append(answerMapBeforeUpdateHooks, answerMapHook)
	case boil.AfterUpdateHook:
		answerMapAfterUpdateHooks = append(answerMapAfterUpdateHooks, answerMapHook)
	case boil.BeforeDeleteHook:
		answerMapBeforeDeleteHooks = append(answerMapBeforeDeleteHooks, answerMapHook)
	case boil.AfterDeleteHook:
		answerMapAfterDeleteHooks = append(answerMapAfterDeleteHooks, answerMapHook)
	case boil.BeforeUpsertHook:
		answerMapBeforeUpsertHooks = append(answerMapBeforeUpsertHooks, answerMapHook)
	case boil.AfterUpsertHook:
		answerMapAfterUpsertHooks = append(answerMapAfterUpsertHooks, answerMapHook)
	}
}

// One returns a single answerMap record from the query.
func (q answerMapQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AnswerMap, error) {
	o := &AnswerMap{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for answer_map")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AnswerMap records from the query.
func (q answerMapQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnswerMapSlice, error) {
	var o []*AnswerMap

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to AnswerMap slice")
	}

	if len(answerMapAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AnswerMap records in the query.
func (q answerMapQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count answer_map rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q answerMapQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if answer_map exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *AnswerMap) User(mods ...qm.QueryMod) answerUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"user_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return AnswerUsers(queryMods...)
}

// MapAnswerEvents retrieves all the answer_event's AnswerEvents with an executor via map_id column.
func (o *AnswerMap) MapAnswerEvents(mods ...qm.QueryMod) answerEventQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"answer_event\".\"map_id\"=?", o.MapID),
	)

	return AnswerEvents(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (answerMapL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnswerMap interface{}, mods queries.Applicator) error {
	var slice []*AnswerMap
	var object *AnswerMap

	if singular {
		object = maybeAnswerMap.(*AnswerMap)
	} else {
		slice = *maybeAnswerMap.(*[]*AnswerMap)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &answerMapR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &answerMapR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`answer_user`),
		qm.WhereIn(`answer_user.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AnswerUser")
	}

	var resultSlice []*AnswerUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AnswerUser")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for answer_user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answer_user")
	}

	if len(answerMapAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &answerUserR{}
		}
		foreign.R.UserAnswerMaps = append(foreign.R.UserAnswerMaps, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &answerUserR{}
				}
				foreign.R.UserAnswerMaps = append(foreign.R.UserAnswerMaps, local)
				break
			}
		}
	}

	return nil
}

// LoadMapAnswerEvents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (answerMapL) LoadMapAnswerEvents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnswerMap interface{}, mods queries.Applicator) error {
	var slice []*AnswerMap
	var object *AnswerMap

	if singular {
		object = maybeAnswerMap.(*AnswerMap)
	} else {
		slice = *maybeAnswerMap.(*[]*AnswerMap)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &answerMapR{}
		}
		args = append(args, object.MapID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &answerMapR{}
			}

			for _, a := range args {
				if a == obj.MapID {
					continue Outer
				}
			}

			args = append(args, obj.MapID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`answer_event`),
		qm.WhereIn(`answer_event.map_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load answer_event")
	}

	var resultSlice []*AnswerEvent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice answer_event")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on answer_event")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answer_event")
	}

	if len(answerEventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.MapAnswerEvents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &answerEventR{}
			}
			foreign.R.Map = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.MapID == foreign.MapID {
				local.R.MapAnswerEvents = append(local.R.MapAnswerEvents, foreign)
				if foreign.R == nil {
					foreign.R = &answerEventR{}
				}
				foreign.R.Map = local
				break
			}
		}
	}

	return nil
}

// SetUser of the answerMap to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAnswerMaps.
func (o *AnswerMap) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AnswerUser) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"answer_map\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, answerMapPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.MapID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &answerMapR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &answerUserR{
			UserAnswerMaps: AnswerMapSlice{o},
		}
	} else {
		related.R.UserAnswerMaps = append(related.R.UserAnswerMaps, o)
	}

	return nil
}

// AddMapAnswerEvents adds the given related objects to the existing relationships
// of the answer_map, optionally inserting them as new records.
// Appends related to o.R.MapAnswerEvents.
// Sets related.R.Map appropriately.
func (o *AnswerMap) AddMapAnswerEvents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AnswerEvent) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MapID = o.MapID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"answer_event\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"map_id"}),
				strmangle.WhereClause("\"", "\"", 2, answerEventPrimaryKeyColumns),
			)
			values := []interface{}{o.MapID, rel.EventID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MapID = o.MapID
		}
	}

	if o.R == nil {
		o.R = &answerMapR{
			MapAnswerEvents: related,
		}
	} else {
		o.R.MapAnswerEvents = append(o.R.MapAnswerEvents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &answerEventR{
				Map: o,
			}
		} else {
			rel.R.Map = o
		}
	}
	return nil
}

// AnswerMaps retrieves all the records using an executor.
func AnswerMaps(mods ...qm.QueryMod) answerMapQuery {
	mods = append(mods, qm.From("\"answer_map\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"answer_map\".*"})
	}

	return answerMapQuery{q}
}

// FindAnswerMap retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnswerMap(ctx context.Context, exec boil.ContextExecutor, mapID string, selectCols ...string) (*AnswerMap, error) {
	answerMapObj := &AnswerMap{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"answer_map\" where \"map_id\"=$1", sel,
	)

	q := queries.Raw(query, mapID)

	err := q.Bind(ctx, exec, answerMapObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from answer_map")
	}

	if err = answerMapObj.doAfterSelectHooks(ctx, exec); err != nil {
		return answerMapObj, err
	}

	return answerMapObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AnswerMap) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no answer_map provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerMapColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	answerMapInsertCacheMut.RLock()
	cache, cached := answerMapInsertCache[key]
	answerMapInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			answerMapAllColumns,
			answerMapColumnsWithDefault,
			answerMapColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(answerMapType, answerMapMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(answerMapType, answerMapMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"answer_map\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"answer_map\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into answer_map")
	}

	if !cached {
		answerMapInsertCacheMut.Lock()
		answerMapInsertCache[key] = cache
		answerMapInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AnswerMap.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AnswerMap) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	answerMapUpdateCacheMut.RLock()
	cache, cached := answerMapUpdateCache[key]
	answerMapUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			answerMapAllColumns,
			answerMapPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update answer_map, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"answer_map\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, answerMapPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(answerMapType, answerMapMapping, append(wl, answerMapPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update answer_map row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for answer_map")
	}

	if !cached {
		answerMapUpdateCacheMut.Lock()
		answerMapUpdateCache[key] = cache
		answerMapUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q answerMapQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for answer_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for answer_map")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnswerMapSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"answer_map\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, answerMapPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in answerMap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all answerMap")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AnswerMap) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no answer_map provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerMapColumnsWithDefault, o)

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

	answerMapUpsertCacheMut.RLock()
	cache, cached := answerMapUpsertCache[key]
	answerMapUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			answerMapAllColumns,
			answerMapColumnsWithDefault,
			answerMapColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			answerMapAllColumns,
			answerMapPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert answer_map, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(answerMapPrimaryKeyColumns))
			copy(conflict, answerMapPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"answer_map\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(answerMapType, answerMapMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(answerMapType, answerMapMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert answer_map")
	}

	if !cached {
		answerMapUpsertCacheMut.Lock()
		answerMapUpsertCache[key] = cache
		answerMapUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AnswerMap record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AnswerMap) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no AnswerMap provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), answerMapPrimaryKeyMapping)
	sql := "DELETE FROM \"answer_map\" WHERE \"map_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from answer_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for answer_map")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q answerMapQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no answerMapQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from answer_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for answer_map")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnswerMapSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(answerMapBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"answer_map\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerMapPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from answerMap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for answer_map")
	}

	if len(answerMapAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AnswerMap) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnswerMap(ctx, exec, o.MapID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnswerMapSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnswerMapSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"answer_map\".* FROM \"answer_map\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerMapPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in AnswerMapSlice")
	}

	*o = slice

	return nil
}

// AnswerMapExists checks if the AnswerMap row exists.
func AnswerMapExists(ctx context.Context, exec boil.ContextExecutor, mapID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"answer_map\" where \"map_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, mapID)
	}
	row := exec.QueryRowContext(ctx, sql, mapID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if answer_map exists")
	}

	return exists, nil
}
