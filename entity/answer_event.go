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

// AnswerEvent is an object representing the database table.
type AnswerEvent struct {
	EventID        string      `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	EventType      string      `boil:"event_type" json:"event_type" toml:"event_type" yaml:"event_type"`
	EventTimestamp time.Time   `boil:"event_timestamp" json:"event_timestamp" toml:"event_timestamp" yaml:"event_timestamp"`
	MapID          string      `boil:"map_id" json:"map_id" toml:"map_id" yaml:"map_id"`
	AnswerValue    null.String `boil:"answer_value" json:"answer_value,omitempty" toml:"answer_value" yaml:"answer_value,omitempty"`

	R *answerEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L answerEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnswerEventColumns = struct {
	EventID        string
	EventType      string
	EventTimestamp string
	MapID          string
	AnswerValue    string
}{
	EventID:        "event_id",
	EventType:      "event_type",
	EventTimestamp: "event_timestamp",
	MapID:          "map_id",
	AnswerValue:    "answer_value",
}

var AnswerEventTableColumns = struct {
	EventID        string
	EventType      string
	EventTimestamp string
	MapID          string
	AnswerValue    string
}{
	EventID:        "answer_event.event_id",
	EventType:      "answer_event.event_type",
	EventTimestamp: "answer_event.event_timestamp",
	MapID:          "answer_event.map_id",
	AnswerValue:    "answer_event.answer_value",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var AnswerEventWhere = struct {
	EventID        whereHelperstring
	EventType      whereHelperstring
	EventTimestamp whereHelpertime_Time
	MapID          whereHelperstring
	AnswerValue    whereHelpernull_String
}{
	EventID:        whereHelperstring{field: "\"answer_event\".\"event_id\""},
	EventType:      whereHelperstring{field: "\"answer_event\".\"event_type\""},
	EventTimestamp: whereHelpertime_Time{field: "\"answer_event\".\"event_timestamp\""},
	MapID:          whereHelperstring{field: "\"answer_event\".\"map_id\""},
	AnswerValue:    whereHelpernull_String{field: "\"answer_event\".\"answer_value\""},
}

// AnswerEventRels is where relationship names are stored.
var AnswerEventRels = struct {
	Map string
}{
	Map: "Map",
}

// answerEventR is where relationships are stored.
type answerEventR struct {
	Map *AnswerMap `boil:"Map" json:"Map" toml:"Map" yaml:"Map"`
}

// NewStruct creates a new relationship struct
func (*answerEventR) NewStruct() *answerEventR {
	return &answerEventR{}
}

// answerEventL is where Load methods for each relationship are stored.
type answerEventL struct{}

var (
	answerEventAllColumns            = []string{"event_id", "event_type", "event_timestamp", "map_id", "answer_value"}
	answerEventColumnsWithoutDefault = []string{"event_id", "event_type", "event_timestamp", "map_id"}
	answerEventColumnsWithDefault    = []string{"answer_value"}
	answerEventPrimaryKeyColumns     = []string{"event_id"}
	answerEventGeneratedColumns      = []string{}
)

type (
	// AnswerEventSlice is an alias for a slice of pointers to AnswerEvent.
	// This should almost always be used instead of []AnswerEvent.
	AnswerEventSlice []*AnswerEvent
	// AnswerEventHook is the signature for custom AnswerEvent hook methods
	AnswerEventHook func(context.Context, boil.ContextExecutor, *AnswerEvent) error

	answerEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	answerEventType                 = reflect.TypeOf(&AnswerEvent{})
	answerEventMapping              = queries.MakeStructMapping(answerEventType)
	answerEventPrimaryKeyMapping, _ = queries.BindMapping(answerEventType, answerEventMapping, answerEventPrimaryKeyColumns)
	answerEventInsertCacheMut       sync.RWMutex
	answerEventInsertCache          = make(map[string]insertCache)
	answerEventUpdateCacheMut       sync.RWMutex
	answerEventUpdateCache          = make(map[string]updateCache)
	answerEventUpsertCacheMut       sync.RWMutex
	answerEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var answerEventAfterSelectHooks []AnswerEventHook

var answerEventBeforeInsertHooks []AnswerEventHook
var answerEventAfterInsertHooks []AnswerEventHook

var answerEventBeforeUpdateHooks []AnswerEventHook
var answerEventAfterUpdateHooks []AnswerEventHook

var answerEventBeforeDeleteHooks []AnswerEventHook
var answerEventAfterDeleteHooks []AnswerEventHook

var answerEventBeforeUpsertHooks []AnswerEventHook
var answerEventAfterUpsertHooks []AnswerEventHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AnswerEvent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AnswerEvent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AnswerEvent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AnswerEvent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AnswerEvent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AnswerEvent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AnswerEvent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AnswerEvent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AnswerEvent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerEventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnswerEventHook registers your hook function for all future operations.
func AddAnswerEventHook(hookPoint boil.HookPoint, answerEventHook AnswerEventHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		answerEventAfterSelectHooks = append(answerEventAfterSelectHooks, answerEventHook)
	case boil.BeforeInsertHook:
		answerEventBeforeInsertHooks = append(answerEventBeforeInsertHooks, answerEventHook)
	case boil.AfterInsertHook:
		answerEventAfterInsertHooks = append(answerEventAfterInsertHooks, answerEventHook)
	case boil.BeforeUpdateHook:
		answerEventBeforeUpdateHooks = append(answerEventBeforeUpdateHooks, answerEventHook)
	case boil.AfterUpdateHook:
		answerEventAfterUpdateHooks = append(answerEventAfterUpdateHooks, answerEventHook)
	case boil.BeforeDeleteHook:
		answerEventBeforeDeleteHooks = append(answerEventBeforeDeleteHooks, answerEventHook)
	case boil.AfterDeleteHook:
		answerEventAfterDeleteHooks = append(answerEventAfterDeleteHooks, answerEventHook)
	case boil.BeforeUpsertHook:
		answerEventBeforeUpsertHooks = append(answerEventBeforeUpsertHooks, answerEventHook)
	case boil.AfterUpsertHook:
		answerEventAfterUpsertHooks = append(answerEventAfterUpsertHooks, answerEventHook)
	}
}

// One returns a single answerEvent record from the query.
func (q answerEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AnswerEvent, error) {
	o := &AnswerEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for answer_event")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AnswerEvent records from the query.
func (q answerEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnswerEventSlice, error) {
	var o []*AnswerEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to AnswerEvent slice")
	}

	if len(answerEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AnswerEvent records in the query.
func (q answerEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count answer_event rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q answerEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if answer_event exists")
	}

	return count > 0, nil
}

// Map pointed to by the foreign key.
func (o *AnswerEvent) Map(mods ...qm.QueryMod) answerMapQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"map_id\" = ?", o.MapID),
	}

	queryMods = append(queryMods, mods...)

	return AnswerMaps(queryMods...)
}

// LoadMap allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (answerEventL) LoadMap(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnswerEvent interface{}, mods queries.Applicator) error {
	var slice []*AnswerEvent
	var object *AnswerEvent

	if singular {
		object = maybeAnswerEvent.(*AnswerEvent)
	} else {
		slice = *maybeAnswerEvent.(*[]*AnswerEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &answerEventR{}
		}
		args = append(args, object.MapID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &answerEventR{}
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
		qm.From(`answer_map`),
		qm.WhereIn(`answer_map.map_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AnswerMap")
	}

	var resultSlice []*AnswerMap
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AnswerMap")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for answer_map")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answer_map")
	}

	if len(answerEventAfterSelectHooks) != 0 {
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
		object.R.Map = foreign
		if foreign.R == nil {
			foreign.R = &answerMapR{}
		}
		foreign.R.MapAnswerEvents = append(foreign.R.MapAnswerEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MapID == foreign.MapID {
				local.R.Map = foreign
				if foreign.R == nil {
					foreign.R = &answerMapR{}
				}
				foreign.R.MapAnswerEvents = append(foreign.R.MapAnswerEvents, local)
				break
			}
		}
	}

	return nil
}

// SetMap of the answerEvent to the related item.
// Sets o.R.Map to related.
// Adds o to related.R.MapAnswerEvents.
func (o *AnswerEvent) SetMap(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AnswerMap) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"answer_event\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"map_id"}),
		strmangle.WhereClause("\"", "\"", 2, answerEventPrimaryKeyColumns),
	)
	values := []interface{}{related.MapID, o.EventID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MapID = related.MapID
	if o.R == nil {
		o.R = &answerEventR{
			Map: related,
		}
	} else {
		o.R.Map = related
	}

	if related.R == nil {
		related.R = &answerMapR{
			MapAnswerEvents: AnswerEventSlice{o},
		}
	} else {
		related.R.MapAnswerEvents = append(related.R.MapAnswerEvents, o)
	}

	return nil
}

// AnswerEvents retrieves all the records using an executor.
func AnswerEvents(mods ...qm.QueryMod) answerEventQuery {
	mods = append(mods, qm.From("\"answer_event\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"answer_event\".*"})
	}

	return answerEventQuery{q}
}

// FindAnswerEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnswerEvent(ctx context.Context, exec boil.ContextExecutor, eventID string, selectCols ...string) (*AnswerEvent, error) {
	answerEventObj := &AnswerEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"answer_event\" where \"event_id\"=$1", sel,
	)

	q := queries.Raw(query, eventID)

	err := q.Bind(ctx, exec, answerEventObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from answer_event")
	}

	if err = answerEventObj.doAfterSelectHooks(ctx, exec); err != nil {
		return answerEventObj, err
	}

	return answerEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AnswerEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no answer_event provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	answerEventInsertCacheMut.RLock()
	cache, cached := answerEventInsertCache[key]
	answerEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			answerEventAllColumns,
			answerEventColumnsWithDefault,
			answerEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(answerEventType, answerEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(answerEventType, answerEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"answer_event\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"answer_event\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into answer_event")
	}

	if !cached {
		answerEventInsertCacheMut.Lock()
		answerEventInsertCache[key] = cache
		answerEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AnswerEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AnswerEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	answerEventUpdateCacheMut.RLock()
	cache, cached := answerEventUpdateCache[key]
	answerEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			answerEventAllColumns,
			answerEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update answer_event, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"answer_event\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, answerEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(answerEventType, answerEventMapping, append(wl, answerEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update answer_event row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for answer_event")
	}

	if !cached {
		answerEventUpdateCacheMut.Lock()
		answerEventUpdateCache[key] = cache
		answerEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q answerEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for answer_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for answer_event")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnswerEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"answer_event\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, answerEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in answerEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all answerEvent")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AnswerEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no answer_event provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerEventColumnsWithDefault, o)

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

	answerEventUpsertCacheMut.RLock()
	cache, cached := answerEventUpsertCache[key]
	answerEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			answerEventAllColumns,
			answerEventColumnsWithDefault,
			answerEventColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			answerEventAllColumns,
			answerEventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert answer_event, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(answerEventPrimaryKeyColumns))
			copy(conflict, answerEventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"answer_event\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(answerEventType, answerEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(answerEventType, answerEventMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert answer_event")
	}

	if !cached {
		answerEventUpsertCacheMut.Lock()
		answerEventUpsertCache[key] = cache
		answerEventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AnswerEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AnswerEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no AnswerEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), answerEventPrimaryKeyMapping)
	sql := "DELETE FROM \"answer_event\" WHERE \"event_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from answer_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for answer_event")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q answerEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no answerEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from answer_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for answer_event")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnswerEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(answerEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"answer_event\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from answerEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for answer_event")
	}

	if len(answerEventAfterDeleteHooks) != 0 {
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
func (o *AnswerEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnswerEvent(ctx, exec, o.EventID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnswerEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnswerEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"answer_event\".* FROM \"answer_event\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in AnswerEventSlice")
	}

	*o = slice

	return nil
}

// AnswerEventExists checks if the AnswerEvent row exists.
func AnswerEventExists(ctx context.Context, exec boil.ContextExecutor, eventID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"answer_event\" where \"event_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, eventID)
	}
	row := exec.QueryRowContext(ctx, sql, eventID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if answer_event exists")
	}

	return exists, nil
}
