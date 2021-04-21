// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repo

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

// LineRevision is an object representing the database table.
type LineRevision struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	ReviewerID int         `boil:"reviewer_id" json:"reviewer_id" toml:"reviewer_id" yaml:"reviewer_id"`
	LineID     int         `boil:"line_id" json:"line_id" toml:"line_id" yaml:"line_id"`
	LineText   null.String `boil:"line_text" json:"line_text,omitempty" toml:"line_text" yaml:"line_text,omitempty"`
	CreatedAt  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt  null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *lineRevisionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lineRevisionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LineRevisionColumns = struct {
	ID         string
	ReviewerID string
	LineID     string
	LineText   string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}{
	ID:         "id",
	ReviewerID: "reviewer_id",
	LineID:     "line_id",
	LineText:   "line_text",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// Generated where

var LineRevisionWhere = struct {
	ID         whereHelperint
	ReviewerID whereHelperint
	LineID     whereHelperint
	LineText   whereHelpernull_String
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	DeletedAt  whereHelpernull_Time
}{
	ID:         whereHelperint{field: "\"line_revision\".\"id\""},
	ReviewerID: whereHelperint{field: "\"line_revision\".\"reviewer_id\""},
	LineID:     whereHelperint{field: "\"line_revision\".\"line_id\""},
	LineText:   whereHelpernull_String{field: "\"line_revision\".\"line_text\""},
	CreatedAt:  whereHelpertime_Time{field: "\"line_revision\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"line_revision\".\"updated_at\""},
	DeletedAt:  whereHelpernull_Time{field: "\"line_revision\".\"deleted_at\""},
}

// LineRevisionRels is where relationship names are stored.
var LineRevisionRels = struct {
	Line     string
	Reviewer string
}{
	Line:     "Line",
	Reviewer: "Reviewer",
}

// lineRevisionR is where relationships are stored.
type lineRevisionR struct {
	Line     *Line `boil:"Line" json:"Line" toml:"Line" yaml:"Line"`
	Reviewer *User `boil:"Reviewer" json:"Reviewer" toml:"Reviewer" yaml:"Reviewer"`
}

// NewStruct creates a new relationship struct
func (*lineRevisionR) NewStruct() *lineRevisionR {
	return &lineRevisionR{}
}

// lineRevisionL is where Load methods for each relationship are stored.
type lineRevisionL struct{}

var (
	lineRevisionAllColumns            = []string{"id", "reviewer_id", "line_id", "line_text", "created_at", "updated_at", "deleted_at"}
	lineRevisionColumnsWithoutDefault = []string{"reviewer_id", "line_id", "line_text", "deleted_at"}
	lineRevisionColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	lineRevisionPrimaryKeyColumns     = []string{"id"}
)

type (
	// LineRevisionSlice is an alias for a slice of pointers to LineRevision.
	// This should generally be used opposed to []LineRevision.
	LineRevisionSlice []*LineRevision
	// LineRevisionHook is the signature for custom LineRevision hook methods
	LineRevisionHook func(context.Context, boil.ContextExecutor, *LineRevision) error

	lineRevisionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lineRevisionType                 = reflect.TypeOf(&LineRevision{})
	lineRevisionMapping              = queries.MakeStructMapping(lineRevisionType)
	lineRevisionPrimaryKeyMapping, _ = queries.BindMapping(lineRevisionType, lineRevisionMapping, lineRevisionPrimaryKeyColumns)
	lineRevisionInsertCacheMut       sync.RWMutex
	lineRevisionInsertCache          = make(map[string]insertCache)
	lineRevisionUpdateCacheMut       sync.RWMutex
	lineRevisionUpdateCache          = make(map[string]updateCache)
	lineRevisionUpsertCacheMut       sync.RWMutex
	lineRevisionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var lineRevisionBeforeInsertHooks []LineRevisionHook
var lineRevisionBeforeUpdateHooks []LineRevisionHook
var lineRevisionBeforeDeleteHooks []LineRevisionHook
var lineRevisionBeforeUpsertHooks []LineRevisionHook

var lineRevisionAfterInsertHooks []LineRevisionHook
var lineRevisionAfterSelectHooks []LineRevisionHook
var lineRevisionAfterUpdateHooks []LineRevisionHook
var lineRevisionAfterDeleteHooks []LineRevisionHook
var lineRevisionAfterUpsertHooks []LineRevisionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LineRevision) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LineRevision) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LineRevision) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LineRevision) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LineRevision) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LineRevision) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LineRevision) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LineRevision) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LineRevision) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lineRevisionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLineRevisionHook registers your hook function for all future operations.
func AddLineRevisionHook(hookPoint boil.HookPoint, lineRevisionHook LineRevisionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		lineRevisionBeforeInsertHooks = append(lineRevisionBeforeInsertHooks, lineRevisionHook)
	case boil.BeforeUpdateHook:
		lineRevisionBeforeUpdateHooks = append(lineRevisionBeforeUpdateHooks, lineRevisionHook)
	case boil.BeforeDeleteHook:
		lineRevisionBeforeDeleteHooks = append(lineRevisionBeforeDeleteHooks, lineRevisionHook)
	case boil.BeforeUpsertHook:
		lineRevisionBeforeUpsertHooks = append(lineRevisionBeforeUpsertHooks, lineRevisionHook)
	case boil.AfterInsertHook:
		lineRevisionAfterInsertHooks = append(lineRevisionAfterInsertHooks, lineRevisionHook)
	case boil.AfterSelectHook:
		lineRevisionAfterSelectHooks = append(lineRevisionAfterSelectHooks, lineRevisionHook)
	case boil.AfterUpdateHook:
		lineRevisionAfterUpdateHooks = append(lineRevisionAfterUpdateHooks, lineRevisionHook)
	case boil.AfterDeleteHook:
		lineRevisionAfterDeleteHooks = append(lineRevisionAfterDeleteHooks, lineRevisionHook)
	case boil.AfterUpsertHook:
		lineRevisionAfterUpsertHooks = append(lineRevisionAfterUpsertHooks, lineRevisionHook)
	}
}

// One returns a single lineRevision record from the query.
func (q lineRevisionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LineRevision, error) {
	o := &LineRevision{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repo: failed to execute a one query for line_revision")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LineRevision records from the query.
func (q lineRevisionQuery) All(ctx context.Context, exec boil.ContextExecutor) (LineRevisionSlice, error) {
	var o []*LineRevision

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repo: failed to assign all query results to LineRevision slice")
	}

	if len(lineRevisionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LineRevision records in the query.
func (q lineRevisionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repo: failed to count line_revision rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lineRevisionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repo: failed to check if line_revision exists")
	}

	return count > 0, nil
}

// Line pointed to by the foreign key.
func (o *LineRevision) Line(mods ...qm.QueryMod) lineQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LineID),
	}

	queryMods = append(queryMods, mods...)

	query := Lines(queryMods...)
	queries.SetFrom(query.Query, "\"line\"")

	return query
}

// Reviewer pointed to by the foreign key.
func (o *LineRevision) Reviewer(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ReviewerID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"user\"")

	return query
}

// LoadLine allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lineRevisionL) LoadLine(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLineRevision interface{}, mods queries.Applicator) error {
	var slice []*LineRevision
	var object *LineRevision

	if singular {
		object = maybeLineRevision.(*LineRevision)
	} else {
		slice = *maybeLineRevision.(*[]*LineRevision)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lineRevisionR{}
		}
		args = append(args, object.LineID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lineRevisionR{}
			}

			for _, a := range args {
				if a == obj.LineID {
					continue Outer
				}
			}

			args = append(args, obj.LineID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`line`),
		qm.WhereIn(`line.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Line")
	}

	var resultSlice []*Line
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Line")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for line")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for line")
	}

	if len(lineRevisionAfterSelectHooks) != 0 {
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
		object.R.Line = foreign
		if foreign.R == nil {
			foreign.R = &lineR{}
		}
		foreign.R.LineRevisions = append(foreign.R.LineRevisions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LineID == foreign.ID {
				local.R.Line = foreign
				if foreign.R == nil {
					foreign.R = &lineR{}
				}
				foreign.R.LineRevisions = append(foreign.R.LineRevisions, local)
				break
			}
		}
	}

	return nil
}

// LoadReviewer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lineRevisionL) LoadReviewer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLineRevision interface{}, mods queries.Applicator) error {
	var slice []*LineRevision
	var object *LineRevision

	if singular {
		object = maybeLineRevision.(*LineRevision)
	} else {
		slice = *maybeLineRevision.(*[]*LineRevision)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lineRevisionR{}
		}
		args = append(args, object.ReviewerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lineRevisionR{}
			}

			for _, a := range args {
				if a == obj.ReviewerID {
					continue Outer
				}
			}

			args = append(args, obj.ReviewerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(lineRevisionAfterSelectHooks) != 0 {
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
		object.R.Reviewer = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.ReviewerLineRevisions = append(foreign.R.ReviewerLineRevisions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ReviewerID == foreign.ID {
				local.R.Reviewer = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.ReviewerLineRevisions = append(foreign.R.ReviewerLineRevisions, local)
				break
			}
		}
	}

	return nil
}

// SetLine of the lineRevision to the related item.
// Sets o.R.Line to related.
// Adds o to related.R.LineRevisions.
func (o *LineRevision) SetLine(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Line) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"line_revision\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"line_id"}),
		strmangle.WhereClause("\"", "\"", 2, lineRevisionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LineID = related.ID
	if o.R == nil {
		o.R = &lineRevisionR{
			Line: related,
		}
	} else {
		o.R.Line = related
	}

	if related.R == nil {
		related.R = &lineR{
			LineRevisions: LineRevisionSlice{o},
		}
	} else {
		related.R.LineRevisions = append(related.R.LineRevisions, o)
	}

	return nil
}

// SetReviewer of the lineRevision to the related item.
// Sets o.R.Reviewer to related.
// Adds o to related.R.ReviewerLineRevisions.
func (o *LineRevision) SetReviewer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"line_revision\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"reviewer_id"}),
		strmangle.WhereClause("\"", "\"", 2, lineRevisionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ReviewerID = related.ID
	if o.R == nil {
		o.R = &lineRevisionR{
			Reviewer: related,
		}
	} else {
		o.R.Reviewer = related
	}

	if related.R == nil {
		related.R = &userR{
			ReviewerLineRevisions: LineRevisionSlice{o},
		}
	} else {
		related.R.ReviewerLineRevisions = append(related.R.ReviewerLineRevisions, o)
	}

	return nil
}

// LineRevisions retrieves all the records using an executor.
func LineRevisions(mods ...qm.QueryMod) lineRevisionQuery {
	mods = append(mods, qm.From("\"line_revision\""))
	return lineRevisionQuery{NewQuery(mods...)}
}

// FindLineRevision retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLineRevision(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*LineRevision, error) {
	lineRevisionObj := &LineRevision{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"line_revision\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lineRevisionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repo: unable to select from line_revision")
	}

	return lineRevisionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LineRevision) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("repo: no line_revision provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(lineRevisionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lineRevisionInsertCacheMut.RLock()
	cache, cached := lineRevisionInsertCache[key]
	lineRevisionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lineRevisionAllColumns,
			lineRevisionColumnsWithDefault,
			lineRevisionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lineRevisionType, lineRevisionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lineRevisionType, lineRevisionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"line_revision\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"line_revision\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "repo: unable to insert into line_revision")
	}

	if !cached {
		lineRevisionInsertCacheMut.Lock()
		lineRevisionInsertCache[key] = cache
		lineRevisionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LineRevision.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LineRevision) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	lineRevisionUpdateCacheMut.RLock()
	cache, cached := lineRevisionUpdateCache[key]
	lineRevisionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lineRevisionAllColumns,
			lineRevisionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("repo: unable to update line_revision, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"line_revision\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lineRevisionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lineRevisionType, lineRevisionMapping, append(wl, lineRevisionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "repo: unable to update line_revision row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repo: failed to get rows affected by update for line_revision")
	}

	if !cached {
		lineRevisionUpdateCacheMut.Lock()
		lineRevisionUpdateCache[key] = cache
		lineRevisionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q lineRevisionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to update all for line_revision")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to retrieve rows affected for line_revision")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LineRevisionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("repo: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lineRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"line_revision\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lineRevisionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to update all in lineRevision slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to retrieve rows affected all in update all lineRevision")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LineRevision) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("repo: no line_revision provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(lineRevisionColumnsWithDefault, o)

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

	lineRevisionUpsertCacheMut.RLock()
	cache, cached := lineRevisionUpsertCache[key]
	lineRevisionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lineRevisionAllColumns,
			lineRevisionColumnsWithDefault,
			lineRevisionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lineRevisionAllColumns,
			lineRevisionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("repo: unable to upsert line_revision, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lineRevisionPrimaryKeyColumns))
			copy(conflict, lineRevisionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"line_revision\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lineRevisionType, lineRevisionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lineRevisionType, lineRevisionMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "repo: unable to upsert line_revision")
	}

	if !cached {
		lineRevisionUpsertCacheMut.Lock()
		lineRevisionUpsertCache[key] = cache
		lineRevisionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single LineRevision record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LineRevision) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("repo: no LineRevision provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lineRevisionPrimaryKeyMapping)
	sql := "DELETE FROM \"line_revision\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to delete from line_revision")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repo: failed to get rows affected by delete for line_revision")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lineRevisionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("repo: no lineRevisionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to delete all from line_revision")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repo: failed to get rows affected by deleteall for line_revision")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LineRevisionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(lineRevisionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lineRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"line_revision\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lineRevisionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repo: unable to delete all from lineRevision slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repo: failed to get rows affected by deleteall for line_revision")
	}

	if len(lineRevisionAfterDeleteHooks) != 0 {
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
func (o *LineRevision) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLineRevision(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LineRevisionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LineRevisionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lineRevisionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"line_revision\".* FROM \"line_revision\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lineRevisionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "repo: unable to reload all in LineRevisionSlice")
	}

	*o = slice

	return nil
}

// LineRevisionExists checks if the LineRevision row exists.
func LineRevisionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"line_revision\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "repo: unable to check if line_revision exists")
	}

	return exists, nil
}
