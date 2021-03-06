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

func testMigMigrations(t *testing.T) {
	t.Parallel()

	query := MigMigrations(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testMigMigrationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = migMigration.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMigMigrationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = MigMigrations(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMigMigrationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := MigMigrationSlice{migMigration}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testMigMigrationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := MigMigrationExists(tx, migMigration.ID)
	if err != nil {
		t.Errorf("Unable to check if MigMigration exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MigMigrationExistsG to return true, but got false.")
	}
}
func testMigMigrationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	migMigrationFound, err := FindMigMigration(tx, migMigration.ID)
	if err != nil {
		t.Error(err)
	}

	if migMigrationFound == nil {
		t.Error("want a record, got nil")
	}
}
func testMigMigrationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = MigMigrations(tx).Bind(migMigration); err != nil {
		t.Error(err)
	}
}

func testMigMigrationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := MigMigrations(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMigMigrationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigrationOne := &MigMigration{}
	migMigrationTwo := &MigMigration{}
	if err = randomize.Struct(seed, migMigrationOne, migMigrationDBTypes, false, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, migMigrationTwo, migMigrationDBTypes, false, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigrationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = migMigrationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := MigMigrations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMigMigrationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	migMigrationOne := &MigMigration{}
	migMigrationTwo := &MigMigration{}
	if err = randomize.Struct(seed, migMigrationOne, migMigrationDBTypes, false, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, migMigrationTwo, migMigrationDBTypes, false, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigrationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = migMigrationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func migMigrationBeforeInsertHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationAfterInsertHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationAfterSelectHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationBeforeUpdateHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationAfterUpdateHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationBeforeDeleteHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationAfterDeleteHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationBeforeUpsertHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func migMigrationAfterUpsertHook(e boil.Executor, o *MigMigration) error {
	*o = MigMigration{}
	return nil
}

func testMigMigrationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &MigMigration{}
	o := &MigMigration{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, migMigrationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MigMigration object: %s", err)
	}

	AddMigMigrationHook(boil.BeforeInsertHook, migMigrationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	migMigrationBeforeInsertHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.AfterInsertHook, migMigrationAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	migMigrationAfterInsertHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.AfterSelectHook, migMigrationAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	migMigrationAfterSelectHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.BeforeUpdateHook, migMigrationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	migMigrationBeforeUpdateHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.AfterUpdateHook, migMigrationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	migMigrationAfterUpdateHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.BeforeDeleteHook, migMigrationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	migMigrationBeforeDeleteHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.AfterDeleteHook, migMigrationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	migMigrationAfterDeleteHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.BeforeUpsertHook, migMigrationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	migMigrationBeforeUpsertHooks = []MigMigrationHook{}

	AddMigMigrationHook(boil.AfterUpsertHook, migMigrationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	migMigrationAfterUpsertHooks = []MigMigrationHook{}
}
func testMigMigrationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMigMigrationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx, migMigrationColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMigMigrationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = migMigration.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testMigMigrationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := MigMigrationSlice{migMigration}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testMigMigrationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := MigMigrations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	migMigrationDBTypes = map[string]string{`ID`: `integer`, `IsApplied`: `boolean`, `Tstamp`: `timestamp without time zone`, `VersionID`: `bigint`}
	_                   = bytes.MinRead
)

func testMigMigrationsUpdate(t *testing.T) {
	t.Parallel()

	if len(migMigrationColumns) == len(migMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	if err = migMigration.Update(tx); err != nil {
		t.Error(err)
	}
}

func testMigMigrationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(migMigrationColumns) == len(migMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	migMigration := &MigMigration{}
	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, migMigration, migMigrationDBTypes, true, migMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(migMigrationColumns, migMigrationPrimaryKeyColumns) {
		fields = migMigrationColumns
	} else {
		fields = strmangle.SetComplement(
			migMigrationColumns,
			migMigrationPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(migMigration))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := MigMigrationSlice{migMigration}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testMigMigrationsUpsert(t *testing.T) {
	t.Parallel()

	if len(migMigrationColumns) == len(migMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	migMigration := MigMigration{}
	if err = randomize.Struct(seed, &migMigration, migMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = migMigration.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert MigMigration: %s", err)
	}

	count, err := MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &migMigration, migMigrationDBTypes, false, migMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MigMigration struct: %s", err)
	}

	if err = migMigration.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert MigMigration: %s", err)
	}

	count, err = MigMigrations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
