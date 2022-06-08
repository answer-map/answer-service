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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AnswerUser is an object representing the database table.
type AnswerUser struct {
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	R *answerUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L answerUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnswerUserColumns = struct {
	UserID string
}{
	UserID: "user_id",
}

var AnswerUserTableColumns = struct {
	UserID string
}{
	UserID: "answer_user.user_id",
}

// Generated where

var AnswerUserWhere = struct {
	UserID whereHelperstring
}{
	UserID: whereHelperstring{field: "\"answer_user\".\"user_id\""},
}

// AnswerUserRels is where relationship names are stored.
var AnswerUserRels = struct {
	UserAnswerMaps string
}{
	UserAnswerMaps: "UserAnswerMaps",
}

// answerUserR is where relationships are stored.
type answerUserR struct {
	UserAnswerMaps AnswerMapSlice `boil:"UserAnswerMaps" json:"UserAnswerMaps" toml:"UserAnswerMaps" yaml:"UserAnswerMaps"`
}

// NewStruct creates a new relationship struct
func (*answerUserR) NewStruct() *answerUserR {
	return &answerUserR{}
}

// answerUserL is where Load methods for each relationship are stored.
type answerUserL struct{}

var (
	answerUserAllColumns            = []string{"user_id"}
	answerUserColumnsWithoutDefault = []string{"user_id"}
	answerUserColumnsWithDefault    = []string{}
	answerUserPrimaryKeyColumns     = []string{"user_id"}
	answerUserGeneratedColumns      = []string{}
)

type (
	// AnswerUserSlice is an alias for a slice of pointers to AnswerUser.
	// This should almost always be used instead of []AnswerUser.
	AnswerUserSlice []*AnswerUser
	// AnswerUserHook is the signature for custom AnswerUser hook methods
	AnswerUserHook func(context.Context, boil.ContextExecutor, *AnswerUser) error

	answerUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	answerUserType                 = reflect.TypeOf(&AnswerUser{})
	answerUserMapping              = queries.MakeStructMapping(answerUserType)
	answerUserPrimaryKeyMapping, _ = queries.BindMapping(answerUserType, answerUserMapping, answerUserPrimaryKeyColumns)
	answerUserInsertCacheMut       sync.RWMutex
	answerUserInsertCache          = make(map[string]insertCache)
	answerUserUpdateCacheMut       sync.RWMutex
	answerUserUpdateCache          = make(map[string]updateCache)
	answerUserUpsertCacheMut       sync.RWMutex
	answerUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var answerUserAfterSelectHooks []AnswerUserHook

var answerUserBeforeInsertHooks []AnswerUserHook
var answerUserAfterInsertHooks []AnswerUserHook

var answerUserBeforeUpdateHooks []AnswerUserHook
var answerUserAfterUpdateHooks []AnswerUserHook

var answerUserBeforeDeleteHooks []AnswerUserHook
var answerUserAfterDeleteHooks []AnswerUserHook

var answerUserBeforeUpsertHooks []AnswerUserHook
var answerUserAfterUpsertHooks []AnswerUserHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AnswerUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AnswerUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AnswerUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AnswerUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AnswerUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AnswerUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AnswerUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AnswerUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AnswerUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnswerUserHook registers your hook function for all future operations.
func AddAnswerUserHook(hookPoint boil.HookPoint, answerUserHook AnswerUserHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		answerUserAfterSelectHooks = append(answerUserAfterSelectHooks, answerUserHook)
	case boil.BeforeInsertHook:
		answerUserBeforeInsertHooks = append(answerUserBeforeInsertHooks, answerUserHook)
	case boil.AfterInsertHook:
		answerUserAfterInsertHooks = append(answerUserAfterInsertHooks, answerUserHook)
	case boil.BeforeUpdateHook:
		answerUserBeforeUpdateHooks = append(answerUserBeforeUpdateHooks, answerUserHook)
	case boil.AfterUpdateHook:
		answerUserAfterUpdateHooks = append(answerUserAfterUpdateHooks, answerUserHook)
	case boil.BeforeDeleteHook:
		answerUserBeforeDeleteHooks = append(answerUserBeforeDeleteHooks, answerUserHook)
	case boil.AfterDeleteHook:
		answerUserAfterDeleteHooks = append(answerUserAfterDeleteHooks, answerUserHook)
	case boil.BeforeUpsertHook:
		answerUserBeforeUpsertHooks = append(answerUserBeforeUpsertHooks, answerUserHook)
	case boil.AfterUpsertHook:
		answerUserAfterUpsertHooks = append(answerUserAfterUpsertHooks, answerUserHook)
	}
}

// One returns a single answerUser record from the query.
func (q answerUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AnswerUser, error) {
	o := &AnswerUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for answer_user")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AnswerUser records from the query.
func (q answerUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnswerUserSlice, error) {
	var o []*AnswerUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to AnswerUser slice")
	}

	if len(answerUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AnswerUser records in the query.
func (q answerUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count answer_user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q answerUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if answer_user exists")
	}

	return count > 0, nil
}

// UserAnswerMaps retrieves all the answer_map's AnswerMaps with an executor via user_id column.
func (o *AnswerUser) UserAnswerMaps(mods ...qm.QueryMod) answerMapQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"answer_map\".\"user_id\"=?", o.UserID),
	)

	return AnswerMaps(queryMods...)
}

// LoadUserAnswerMaps allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (answerUserL) LoadUserAnswerMaps(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnswerUser interface{}, mods queries.Applicator) error {
	var slice []*AnswerUser
	var object *AnswerUser

	if singular {
		object = maybeAnswerUser.(*AnswerUser)
	} else {
		slice = *maybeAnswerUser.(*[]*AnswerUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &answerUserR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &answerUserR{}
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
		qm.From(`answer_map`),
		qm.WhereIn(`answer_map.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load answer_map")
	}

	var resultSlice []*AnswerMap
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice answer_map")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on answer_map")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answer_map")
	}

	if len(answerMapAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserAnswerMaps = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &answerMapR{}
			}
			foreign.R.User = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UserID == foreign.UserID {
				local.R.UserAnswerMaps = append(local.R.UserAnswerMaps, foreign)
				if foreign.R == nil {
					foreign.R = &answerMapR{}
				}
				foreign.R.User = local
				break
			}
		}
	}

	return nil
}

// AddUserAnswerMaps adds the given related objects to the existing relationships
// of the answer_user, optionally inserting them as new records.
// Appends related to o.R.UserAnswerMaps.
// Sets related.R.User appropriately.
func (o *AnswerUser) AddUserAnswerMaps(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AnswerMap) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserID = o.UserID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"answer_map\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
				strmangle.WhereClause("\"", "\"", 2, answerMapPrimaryKeyColumns),
			)
			values := []interface{}{o.UserID, rel.MapID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.UserID = o.UserID
		}
	}

	if o.R == nil {
		o.R = &answerUserR{
			UserAnswerMaps: related,
		}
	} else {
		o.R.UserAnswerMaps = append(o.R.UserAnswerMaps, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &answerMapR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// AnswerUsers retrieves all the records using an executor.
func AnswerUsers(mods ...qm.QueryMod) answerUserQuery {
	mods = append(mods, qm.From("\"answer_user\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"answer_user\".*"})
	}

	return answerUserQuery{q}
}

// FindAnswerUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnswerUser(ctx context.Context, exec boil.ContextExecutor, userID string, selectCols ...string) (*AnswerUser, error) {
	answerUserObj := &AnswerUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"answer_user\" where \"user_id\"=$1", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, answerUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from answer_user")
	}

	if err = answerUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return answerUserObj, err
	}

	return answerUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AnswerUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no answer_user provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	answerUserInsertCacheMut.RLock()
	cache, cached := answerUserInsertCache[key]
	answerUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			answerUserAllColumns,
			answerUserColumnsWithDefault,
			answerUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(answerUserType, answerUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(answerUserType, answerUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"answer_user\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"answer_user\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into answer_user")
	}

	if !cached {
		answerUserInsertCacheMut.Lock()
		answerUserInsertCache[key] = cache
		answerUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AnswerUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AnswerUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	answerUserUpdateCacheMut.RLock()
	cache, cached := answerUserUpdateCache[key]
	answerUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			answerUserAllColumns,
			answerUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update answer_user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"answer_user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, answerUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(answerUserType, answerUserMapping, append(wl, answerUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update answer_user row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for answer_user")
	}

	if !cached {
		answerUserUpdateCacheMut.Lock()
		answerUserUpdateCache[key] = cache
		answerUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q answerUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for answer_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for answer_user")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnswerUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"answer_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, answerUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in answerUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all answerUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AnswerUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no answer_user provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(answerUserColumnsWithDefault, o)

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

	answerUserUpsertCacheMut.RLock()
	cache, cached := answerUserUpsertCache[key]
	answerUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			answerUserAllColumns,
			answerUserColumnsWithDefault,
			answerUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			answerUserAllColumns,
			answerUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert answer_user, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(answerUserPrimaryKeyColumns))
			copy(conflict, answerUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"answer_user\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(answerUserType, answerUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(answerUserType, answerUserMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert answer_user")
	}

	if !cached {
		answerUserUpsertCacheMut.Lock()
		answerUserUpsertCache[key] = cache
		answerUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AnswerUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AnswerUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no AnswerUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), answerUserPrimaryKeyMapping)
	sql := "DELETE FROM \"answer_user\" WHERE \"user_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from answer_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for answer_user")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q answerUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no answerUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from answer_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for answer_user")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnswerUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(answerUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"answer_user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from answerUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for answer_user")
	}

	if len(answerUserAfterDeleteHooks) != 0 {
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
func (o *AnswerUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnswerUser(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnswerUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnswerUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"answer_user\".* FROM \"answer_user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in AnswerUserSlice")
	}

	*o = slice

	return nil
}

// AnswerUserExists checks if the AnswerUser row exists.
func AnswerUserExists(ctx context.Context, exec boil.ContextExecutor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"answer_user\" where \"user_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID)
	}
	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if answer_user exists")
	}

	return exists, nil
}
