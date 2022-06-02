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

func testAnswerMaps(t *testing.T) {
	t.Parallel()

	query := AnswerMaps()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAnswerMapsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
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

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnswerMapsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AnswerMaps().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnswerMapsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AnswerMapSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnswerMapsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AnswerMapExists(ctx, tx, o.AnswerKey)
	if err != nil {
		t.Errorf("Unable to check if AnswerMap exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnswerMapExists to return true, but got false.")
	}
}

func testAnswerMapsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	answerMapFound, err := FindAnswerMap(ctx, tx, o.AnswerKey)
	if err != nil {
		t.Error(err)
	}

	if answerMapFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAnswerMapsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AnswerMaps().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAnswerMapsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AnswerMaps().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnswerMapsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	answerMapOne := &AnswerMap{}
	answerMapTwo := &AnswerMap{}
	if err = randomize.Struct(seed, answerMapOne, answerMapDBTypes, false, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}
	if err = randomize.Struct(seed, answerMapTwo, answerMapDBTypes, false, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = answerMapOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = answerMapTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AnswerMaps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnswerMapsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	answerMapOne := &AnswerMap{}
	answerMapTwo := &AnswerMap{}
	if err = randomize.Struct(seed, answerMapOne, answerMapDBTypes, false, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}
	if err = randomize.Struct(seed, answerMapTwo, answerMapDBTypes, false, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = answerMapOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = answerMapTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func answerMapBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func answerMapAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AnswerMap) error {
	*o = AnswerMap{}
	return nil
}

func testAnswerMapsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AnswerMap{}
	o := &AnswerMap{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, answerMapDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AnswerMap object: %s", err)
	}

	AddAnswerMapHook(boil.BeforeInsertHook, answerMapBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	answerMapBeforeInsertHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.AfterInsertHook, answerMapAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	answerMapAfterInsertHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.AfterSelectHook, answerMapAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	answerMapAfterSelectHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.BeforeUpdateHook, answerMapBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	answerMapBeforeUpdateHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.AfterUpdateHook, answerMapAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	answerMapAfterUpdateHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.BeforeDeleteHook, answerMapBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	answerMapBeforeDeleteHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.AfterDeleteHook, answerMapAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	answerMapAfterDeleteHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.BeforeUpsertHook, answerMapBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	answerMapBeforeUpsertHooks = []AnswerMapHook{}

	AddAnswerMapHook(boil.AfterUpsertHook, answerMapAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	answerMapAfterUpsertHooks = []AnswerMapHook{}
}

func testAnswerMapsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnswerMapsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(answerMapColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnswerMapToManyAnswerKeyAnswerEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnswerMap
	var b, c AnswerEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, answerEventDBTypes, false, answerEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AnswerKey = a.AnswerKey
	c.AnswerKey = a.AnswerKey

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AnswerKeyAnswerEvents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.AnswerKey == b.AnswerKey {
			bFound = true
		}
		if v.AnswerKey == c.AnswerKey {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AnswerMapSlice{&a}
	if err = a.L.LoadAnswerKeyAnswerEvents(ctx, tx, false, (*[]*AnswerMap)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnswerKeyAnswerEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AnswerKeyAnswerEvents = nil
	if err = a.L.LoadAnswerKeyAnswerEvents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnswerKeyAnswerEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAnswerMapToManyAddOpAnswerKeyAnswerEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnswerMap
	var b, c, d, e AnswerEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, answerMapDBTypes, false, strmangle.SetComplement(answerMapPrimaryKeyColumns, answerMapColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AnswerEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, answerEventDBTypes, false, strmangle.SetComplement(answerEventPrimaryKeyColumns, answerEventColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AnswerEvent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAnswerKeyAnswerEvents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AnswerKey != first.AnswerKey {
			t.Error("foreign key was wrong value", a.AnswerKey, first.AnswerKey)
		}
		if a.AnswerKey != second.AnswerKey {
			t.Error("foreign key was wrong value", a.AnswerKey, second.AnswerKey)
		}

		if first.R.AnswerKeyAnswerMap != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AnswerKeyAnswerMap != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AnswerKeyAnswerEvents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AnswerKeyAnswerEvents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AnswerKeyAnswerEvents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAnswerMapsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
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

func testAnswerMapsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AnswerMapSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAnswerMapsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AnswerMaps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	answerMapDBTypes = map[string]string{`AnswerKey`: `text`, `AnswerValue`: `text`}
	_                = bytes.MinRead
)

func testAnswerMapsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(answerMapPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(answerMapAllColumns) == len(answerMapPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAnswerMapsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(answerMapAllColumns) == len(answerMapPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AnswerMap{}
	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, answerMapDBTypes, true, answerMapPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(answerMapAllColumns, answerMapPrimaryKeyColumns) {
		fields = answerMapAllColumns
	} else {
		fields = strmangle.SetComplement(
			answerMapAllColumns,
			answerMapPrimaryKeyColumns,
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

	slice := AnswerMapSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAnswerMapsUpsert(t *testing.T) {
	t.Parallel()

	if len(answerMapAllColumns) == len(answerMapPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AnswerMap{}
	if err = randomize.Struct(seed, &o, answerMapDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AnswerMap: %s", err)
	}

	count, err := AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, answerMapDBTypes, false, answerMapPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnswerMap struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AnswerMap: %s", err)
	}

	count, err = AnswerMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
