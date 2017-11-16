// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// VotedBallot is an object representing the database table.
type VotedBallot struct {
	ID            string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	VoterEmail    string     `boil:"voter_email" json:"voter_email" toml:"voter_email" yaml:"voter_email"`
	RCV           types.JSON `boil:"rcv" json:"rcv" toml:"rcv" yaml:"rcv"`
	UnexpiredTerm string     `boil:"unexpired_term" json:"unexpired_term" toml:"unexpired_term" yaml:"unexpired_term"`
	VoteFor2      types.JSON `boil:"vote_for_2" json:"vote_for_2" toml:"vote_for_2" yaml:"vote_for_2"`
	BallotIssue   string     `boil:"ballot_issue" json:"ballot_issue" toml:"ballot_issue" yaml:"ballot_issue"`

	R *votedBallotR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L votedBallotL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VotedBallotColumns = struct {
	ID            string
	VoterEmail    string
	RCV           string
	UnexpiredTerm string
	VoteFor2      string
	BallotIssue   string
}{
	ID:            "id",
	VoterEmail:    "voter_email",
	RCV:           "rcv",
	UnexpiredTerm: "unexpired_term",
	VoteFor2:      "vote_for_2",
	BallotIssue:   "ballot_issue",
}

// votedBallotR is where relationships are stored.
type votedBallotR struct {
}

// votedBallotL is where Load methods for each relationship are stored.
type votedBallotL struct{}

var (
	votedBallotColumns               = []string{"id", "voter_email", "rcv", "unexpired_term", "vote_for_2", "ballot_issue"}
	votedBallotColumnsWithoutDefault = []string{"voter_email", "rcv", "unexpired_term", "vote_for_2", "ballot_issue"}
	votedBallotColumnsWithDefault    = []string{"id"}
	votedBallotPrimaryKeyColumns     = []string{"id"}
)

type (
	// VotedBallotSlice is an alias for a slice of pointers to VotedBallot.
	// This should generally be used opposed to []VotedBallot.
	VotedBallotSlice []*VotedBallot
	// VotedBallotHook is the signature for custom VotedBallot hook methods
	VotedBallotHook func(boil.Executor, *VotedBallot) error

	votedBallotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	votedBallotType                 = reflect.TypeOf(&VotedBallot{})
	votedBallotMapping              = queries.MakeStructMapping(votedBallotType)
	votedBallotPrimaryKeyMapping, _ = queries.BindMapping(votedBallotType, votedBallotMapping, votedBallotPrimaryKeyColumns)
	votedBallotInsertCacheMut       sync.RWMutex
	votedBallotInsertCache          = make(map[string]insertCache)
	votedBallotUpdateCacheMut       sync.RWMutex
	votedBallotUpdateCache          = make(map[string]updateCache)
	votedBallotUpsertCacheMut       sync.RWMutex
	votedBallotUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var votedBallotBeforeInsertHooks []VotedBallotHook
var votedBallotBeforeUpdateHooks []VotedBallotHook
var votedBallotBeforeDeleteHooks []VotedBallotHook
var votedBallotBeforeUpsertHooks []VotedBallotHook

var votedBallotAfterInsertHooks []VotedBallotHook
var votedBallotAfterSelectHooks []VotedBallotHook
var votedBallotAfterUpdateHooks []VotedBallotHook
var votedBallotAfterDeleteHooks []VotedBallotHook
var votedBallotAfterUpsertHooks []VotedBallotHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VotedBallot) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VotedBallot) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VotedBallot) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VotedBallot) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VotedBallot) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VotedBallot) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VotedBallot) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VotedBallot) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VotedBallot) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range votedBallotAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVotedBallotHook registers your hook function for all future operations.
func AddVotedBallotHook(hookPoint boil.HookPoint, votedBallotHook VotedBallotHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		votedBallotBeforeInsertHooks = append(votedBallotBeforeInsertHooks, votedBallotHook)
	case boil.BeforeUpdateHook:
		votedBallotBeforeUpdateHooks = append(votedBallotBeforeUpdateHooks, votedBallotHook)
	case boil.BeforeDeleteHook:
		votedBallotBeforeDeleteHooks = append(votedBallotBeforeDeleteHooks, votedBallotHook)
	case boil.BeforeUpsertHook:
		votedBallotBeforeUpsertHooks = append(votedBallotBeforeUpsertHooks, votedBallotHook)
	case boil.AfterInsertHook:
		votedBallotAfterInsertHooks = append(votedBallotAfterInsertHooks, votedBallotHook)
	case boil.AfterSelectHook:
		votedBallotAfterSelectHooks = append(votedBallotAfterSelectHooks, votedBallotHook)
	case boil.AfterUpdateHook:
		votedBallotAfterUpdateHooks = append(votedBallotAfterUpdateHooks, votedBallotHook)
	case boil.AfterDeleteHook:
		votedBallotAfterDeleteHooks = append(votedBallotAfterDeleteHooks, votedBallotHook)
	case boil.AfterUpsertHook:
		votedBallotAfterUpsertHooks = append(votedBallotAfterUpsertHooks, votedBallotHook)
	}
}

// OneP returns a single votedBallot record from the query, and panics on error.
func (q votedBallotQuery) OneP() *VotedBallot {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single votedBallot record from the query.
func (q votedBallotQuery) One() (*VotedBallot, error) {
	o := &VotedBallot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for voted_ballot")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all VotedBallot records from the query, and panics on error.
func (q votedBallotQuery) AllP() VotedBallotSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all VotedBallot records from the query.
func (q votedBallotQuery) All() (VotedBallotSlice, error) {
	var o []*VotedBallot

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VotedBallot slice")
	}

	if len(votedBallotAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all VotedBallot records in the query, and panics on error.
func (q votedBallotQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all VotedBallot records in the query.
func (q votedBallotQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count voted_ballot rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q votedBallotQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q votedBallotQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if voted_ballot exists")
	}

	return count > 0, nil
}

// VotedBallotsG retrieves all records.
func VotedBallotsG(mods ...qm.QueryMod) votedBallotQuery {
	return VotedBallots(boil.GetDB(), mods...)
}

// VotedBallots retrieves all the records using an executor.
func VotedBallots(exec boil.Executor, mods ...qm.QueryMod) votedBallotQuery {
	mods = append(mods, qm.From("\"voted_ballot\""))
	return votedBallotQuery{NewQuery(exec, mods...)}
}

// FindVotedBallotG retrieves a single record by ID.
func FindVotedBallotG(id string, selectCols ...string) (*VotedBallot, error) {
	return FindVotedBallot(boil.GetDB(), id, selectCols...)
}

// FindVotedBallotGP retrieves a single record by ID, and panics on error.
func FindVotedBallotGP(id string, selectCols ...string) *VotedBallot {
	retobj, err := FindVotedBallot(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindVotedBallot retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVotedBallot(exec boil.Executor, id string, selectCols ...string) (*VotedBallot, error) {
	votedBallotObj := &VotedBallot{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"voted_ballot\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(votedBallotObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from voted_ballot")
	}

	return votedBallotObj, nil
}

// FindVotedBallotP retrieves a single record by ID with an executor, and panics on error.
func FindVotedBallotP(exec boil.Executor, id string, selectCols ...string) *VotedBallot {
	retobj, err := FindVotedBallot(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VotedBallot) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *VotedBallot) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *VotedBallot) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *VotedBallot) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no voted_ballot provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(votedBallotColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	votedBallotInsertCacheMut.RLock()
	cache, cached := votedBallotInsertCache[key]
	votedBallotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			votedBallotColumns,
			votedBallotColumnsWithDefault,
			votedBallotColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(votedBallotType, votedBallotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(votedBallotType, votedBallotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"voted_ballot\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"voted_ballot\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
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
		return errors.Wrap(err, "models: unable to insert into voted_ballot")
	}

	if !cached {
		votedBallotInsertCacheMut.Lock()
		votedBallotInsertCache[key] = cache
		votedBallotInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single VotedBallot record. See Update for
// whitelist behavior description.
func (o *VotedBallot) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single VotedBallot record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *VotedBallot) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the VotedBallot, and panics on error.
// See Update for whitelist behavior description.
func (o *VotedBallot) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the VotedBallot.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *VotedBallot) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	votedBallotUpdateCacheMut.RLock()
	cache, cached := votedBallotUpdateCache[key]
	votedBallotUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			votedBallotColumns,
			votedBallotPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update voted_ballot, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"voted_ballot\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, votedBallotPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(votedBallotType, votedBallotMapping, append(wl, votedBallotPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update voted_ballot row")
	}

	if !cached {
		votedBallotUpdateCacheMut.Lock()
		votedBallotUpdateCache[key] = cache
		votedBallotUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q votedBallotQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q votedBallotQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for voted_ballot")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VotedBallotSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o VotedBallotSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o VotedBallotSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VotedBallotSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), votedBallotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"voted_ballot\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, votedBallotPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in votedBallot slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *VotedBallot) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *VotedBallot) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *VotedBallot) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *VotedBallot) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no voted_ballot provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(votedBallotColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	votedBallotUpsertCacheMut.RLock()
	cache, cached := votedBallotUpsertCache[key]
	votedBallotUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			votedBallotColumns,
			votedBallotColumnsWithDefault,
			votedBallotColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			votedBallotColumns,
			votedBallotPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert voted_ballot, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(votedBallotPrimaryKeyColumns))
			copy(conflict, votedBallotPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"voted_ballot\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(votedBallotType, votedBallotMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(votedBallotType, votedBallotMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert voted_ballot")
	}

	if !cached {
		votedBallotUpsertCacheMut.Lock()
		votedBallotUpsertCache[key] = cache
		votedBallotUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single VotedBallot record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *VotedBallot) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single VotedBallot record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *VotedBallot) DeleteG() error {
	if o == nil {
		return errors.New("models: no VotedBallot provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single VotedBallot record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *VotedBallot) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single VotedBallot record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VotedBallot) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no VotedBallot provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), votedBallotPrimaryKeyMapping)
	sql := "DELETE FROM \"voted_ballot\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from voted_ballot")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q votedBallotQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q votedBallotQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no votedBallotQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from voted_ballot")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o VotedBallotSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o VotedBallotSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no VotedBallot slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o VotedBallotSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VotedBallotSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no VotedBallot slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(votedBallotBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), votedBallotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"voted_ballot\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, votedBallotPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from votedBallot slice")
	}

	if len(votedBallotAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *VotedBallot) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *VotedBallot) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *VotedBallot) ReloadG() error {
	if o == nil {
		return errors.New("models: no VotedBallot provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VotedBallot) Reload(exec boil.Executor) error {
	ret, err := FindVotedBallot(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VotedBallotSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *VotedBallotSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VotedBallotSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty VotedBallotSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VotedBallotSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	votedBallots := VotedBallotSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), votedBallotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"voted_ballot\".* FROM \"voted_ballot\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, votedBallotPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&votedBallots)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VotedBallotSlice")
	}

	*o = votedBallots

	return nil
}

// VotedBallotExists checks if the VotedBallot row exists.
func VotedBallotExists(exec boil.Executor, id string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"voted_ballot\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if voted_ballot exists")
	}

	return exists, nil
}

// VotedBallotExistsG checks if the VotedBallot row exists.
func VotedBallotExistsG(id string) (bool, error) {
	return VotedBallotExists(boil.GetDB(), id)
}

// VotedBallotExistsGP checks if the VotedBallot row exists. Panics on error.
func VotedBallotExistsGP(id string) bool {
	e, err := VotedBallotExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// VotedBallotExistsP checks if the VotedBallot row exists. Panics on error.
func VotedBallotExistsP(exec boil.Executor, id string) bool {
	e, err := VotedBallotExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
