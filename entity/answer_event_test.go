// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAnswerEvents(t *testing.T) {
	t.Parallel()

	query := AnswerEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAnswerEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnswerEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AnswerEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnswerEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AnswerEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnswerEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AnswerEventExists(ctx, tx, o.EventID)
	if err != nil {
		t.Errorf("Unable to check if AnswerEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnswerEventExists to return true, but got false.")
	}
}

func testAnswerEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	answerEventFound, err := FindAnswerEvent(ctx, tx, o.EventID)
	if err != nil {
		t.Error(err)
	}

	if answerEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAnswerEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AnswerEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAnswerEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AnswerEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnswerEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	answerEventOne := &AnswerEvent{}
	answerEventTwo := &AnswerEvent{}
	if err = randomize.Struct(seed, answerEventOne, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, answerEventTwo, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = answerEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = answerEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AnswerEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnswerEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	answerEventOne := &AnswerEvent{}
	answerEventTwo := &AnswerEvent{}
	if err = randomize.Struct(seed, answerEventOne, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, answerEventTwo, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = answerEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = answerEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func answerEventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func answerEventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerEvent) error {
	*o = AnswerEvent{}
	return nil
}

func testAnswerEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AnswerEvent{}
	o := &AnswerEvent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, answerEventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AnswerEvent object: %s", err)
	}

	AddAnswerEventHook(boil.BeforeInsertHook, answerEventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	answerEventBeforeInsertHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.AfterInsertHook, answerEventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	answerEventAfterInsertHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.AfterSelectHook, answerEventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	answerEventAfterSelectHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.BeforeUpdateHook, answerEventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	answerEventBeforeUpdateHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.AfterUpdateHook, answerEventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	answerEventAfterUpdateHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.BeforeDeleteHook, answerEventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	answerEventBeforeDeleteHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.AfterDeleteHook, answerEventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	answerEventAfterDeleteHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.BeforeUpsertHook, answerEventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	answerEventBeforeUpsertHooks = []AnswerEventHook{}

	AddAnswerEventHook(boil.AfterUpsertHook, answerEventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	answerEventAfterUpsertHooks = []AnswerEventHook{}
}

func testAnswerEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnswerEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(answerEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnswerEventToOneAnswerMapUsingAnswerKeyAnswerMap(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AnswerEvent
	var foreign AnswerMap

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, answerMapDBTypes, false, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AnswerKey = foreign.AnswerKey
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AnswerKeyAnswerMap().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.AnswerKey != foreign.AnswerKey {
		t.Errorf("want: %v, got %v", foreign.AnswerKey, check.AnswerKey)
	}

	slice := AnswerEventSlice{&local}
	if err = local.L.LoadAnswerKeyAnswerMap(ctx, tx, false, (*[]*AnswerEvent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AnswerKeyAnswerMap == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AnswerKeyAnswerMap = nil
	if err = local.L.LoadAnswerKeyAnswerMap(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AnswerKeyAnswerMap == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnswerEventToOneSetOpAnswerMapUsingAnswerKeyAnswerMap(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnswerEvent
	var b, c AnswerMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, answerEventDBTypes, false, strmangle.SetComplement(answerEventPrimaryKeyColumns, answerEventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, answerMapDBTypes, false, strmangle.SetComplement(answerMapPrimaryKeyColumns, answerMapColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, answerMapDBTypes, false, strmangle.SetComplement(answerMapPrimaryKeyColumns, answerMapColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AnswerMap{&b, &c} {
		err = a.SetAnswerKeyAnswerMap(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AnswerKeyAnswerMap != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AnswerKeyAnswerEvents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AnswerKey != x.AnswerKey {
			t.Error("foreign key was wrong value", a.AnswerKey)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AnswerKey))
		reflect.Indirect(reflect.ValueOf(&a.AnswerKey)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnswerKey != x.AnswerKey {
			t.Error("foreign key was wrong value", a.AnswerKey, x.AnswerKey)
		}
	}
}

func testAnswerEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAnswerEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AnswerEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAnswerEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AnswerEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	answerEventDBTypes = map[string]string{`EventID`: `uuid`, `EventType`: `enum.event_type('create','update','delete')`, `EventTimestamp`: `timestamp with time zone`, `AnswerKey`: `text`, `AnswerValue`: `text`}
	_                  = bytes.MinRead
)

func testAnswerEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(answerEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(answerEventAllColumns) == len(answerEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAnswerEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(answerEventAllColumns) == len(answerEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AnswerEvent{}
	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, answerEventDBTypes, true, answerEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(answerEventAllColumns, answerEventPrimaryKeyColumns) {
		fields = answerEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			answerEventAllColumns,
			answerEventPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AnswerEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAnswerEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(answerEventAllColumns) == len(answerEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AnswerEvent{}
	if err = randomize.Struct(seed, &o, answerEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AnswerEvent: %s", err)
	}

	count, err := AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, answerEventDBTypes, false, answerEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnswerEvent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AnswerEvent: %s", err)
	}

	count, err = AnswerEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
