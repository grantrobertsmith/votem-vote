// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testHumen(t *testing.T) {
	t.Parallel()

	query := Humen(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testHumenDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = human.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHumenQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Humen(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHumenSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := HumanSlice{human}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testHumenExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := HumanExists(tx, human.VotemUhid)
	if err != nil {
		t.Errorf("Unable to check if Human exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HumanExistsG to return true, but got false.")
	}
}
func testHumenFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	humanFound, err := FindHuman(tx, human.VotemUhid)
	if err != nil {
		t.Error(err)
	}

	if humanFound == nil {
		t.Error("want a record, got nil")
	}
}
func testHumenBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Humen(tx).Bind(human); err != nil {
		t.Error(err)
	}
}

func testHumenOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Humen(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHumenAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	humanOne := &Human{}
	humanTwo := &Human{}
	if err = randomize.Struct(seed, humanOne, humanDBTypes, false, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}
	if err = randomize.Struct(seed, humanTwo, humanDBTypes, false, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = humanOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = humanTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Humen(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHumenCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	humanOne := &Human{}
	humanTwo := &Human{}
	if err = randomize.Struct(seed, humanOne, humanDBTypes, false, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}
	if err = randomize.Struct(seed, humanTwo, humanDBTypes, false, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = humanOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = humanTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func humanBeforeInsertHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanAfterInsertHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanAfterSelectHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanBeforeUpdateHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanAfterUpdateHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanBeforeDeleteHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanAfterDeleteHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanBeforeUpsertHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func humanAfterUpsertHook(e boil.Executor, o *Human) error {
	*o = Human{}
	return nil
}

func testHumenHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Human{}
	o := &Human{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, humanDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Human object: %s", err)
	}

	AddHumanHook(boil.BeforeInsertHook, humanBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	humanBeforeInsertHooks = []HumanHook{}

	AddHumanHook(boil.AfterInsertHook, humanAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	humanAfterInsertHooks = []HumanHook{}

	AddHumanHook(boil.AfterSelectHook, humanAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	humanAfterSelectHooks = []HumanHook{}

	AddHumanHook(boil.BeforeUpdateHook, humanBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	humanBeforeUpdateHooks = []HumanHook{}

	AddHumanHook(boil.AfterUpdateHook, humanAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	humanAfterUpdateHooks = []HumanHook{}

	AddHumanHook(boil.BeforeDeleteHook, humanBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	humanBeforeDeleteHooks = []HumanHook{}

	AddHumanHook(boil.AfterDeleteHook, humanAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	humanAfterDeleteHooks = []HumanHook{}

	AddHumanHook(boil.BeforeUpsertHook, humanBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	humanBeforeUpsertHooks = []HumanHook{}

	AddHumanHook(boil.AfterUpsertHook, humanAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	humanAfterUpsertHooks = []HumanHook{}
}
func testHumenInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHumenInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx, humanColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHumanToOneNameRefUsingNameRef(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Human
	var foreign NameRef

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, humanDBTypes, false, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, nameRefDBTypes, false, nameRefColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NameRef struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.NameRefID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.NameRef(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := HumanSlice{&local}
	if err = local.L.LoadNameRef(tx, false, (*[]*Human)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.NameRef == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.NameRef = nil
	if err = local.L.LoadNameRef(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.NameRef == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testHumanToOneSetOpNameRefUsingNameRef(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Human
	var b, c NameRef

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, humanDBTypes, false, strmangle.SetComplement(humanPrimaryKeyColumns, humanColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, nameRefDBTypes, false, strmangle.SetComplement(nameRefPrimaryKeyColumns, nameRefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, nameRefDBTypes, false, strmangle.SetComplement(nameRefPrimaryKeyColumns, nameRefColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*NameRef{&b, &c} {
		err = a.SetNameRef(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.NameRef != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Humen[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.NameRefID != x.ID {
			t.Error("foreign key was wrong value", a.NameRefID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.NameRefID))
		reflect.Indirect(reflect.ValueOf(&a.NameRefID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.NameRefID != x.ID {
			t.Error("foreign key was wrong value", a.NameRefID, x.ID)
		}
	}
}
func testHumenReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = human.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testHumenReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := HumanSlice{human}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testHumenSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Humen(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	humanDBTypes = map[string]string{`NameRefID`: `uuid`, `PasswordHash`: `text`, `PrimaryEmail`: `text`, `PrimaryPhone`: `text`, `Username`: `text`, `VotemUhid`: `uuid`}
	_            = bytes.MinRead
)

func testHumenUpdate(t *testing.T) {
	t.Parallel()

	if len(humanColumns) == len(humanPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	if err = human.Update(tx); err != nil {
		t.Error(err)
	}
}

func testHumenSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(humanColumns) == len(humanPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	human := &Human{}
	if err = randomize.Struct(seed, human, humanDBTypes, true, humanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, human, humanDBTypes, true, humanPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(humanColumns, humanPrimaryKeyColumns) {
		fields = humanColumns
	} else {
		fields = strmangle.SetComplement(
			humanColumns,
			humanPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(human))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := HumanSlice{human}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testHumenUpsert(t *testing.T) {
	t.Parallel()

	if len(humanColumns) == len(humanPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	human := Human{}
	if err = randomize.Struct(seed, &human, humanDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = human.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Human: %s", err)
	}

	count, err := Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &human, humanDBTypes, false, humanPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Human struct: %s", err)
	}

	if err = human.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Human: %s", err)
	}

	count, err = Humen(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}