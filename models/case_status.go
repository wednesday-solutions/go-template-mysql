// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// CaseStatus is an object representing the database table.
type CaseStatus struct {
	ID               int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IncidentReportID int       `boil:"incident_report_id" json:"incident_report_id" toml:"incident_report_id" yaml:"incident_report_id"`
	Status           string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	StatusRemarks    string    `boil:"status_remarks" json:"status_remarks" toml:"status_remarks" yaml:"status_remarks"`
	UserID           int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt        null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *caseStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L caseStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CaseStatusColumns = struct {
	ID               string
	IncidentReportID string
	Status           string
	StatusRemarks    string
	UserID           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "id",
	IncidentReportID: "incident_report_id",
	Status:           "status",
	StatusRemarks:    "status_remarks",
	UserID:           "user_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

var CaseStatusTableColumns = struct {
	ID               string
	IncidentReportID string
	Status           string
	StatusRemarks    string
	UserID           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "case_status.id",
	IncidentReportID: "case_status.incident_report_id",
	Status:           "case_status.status",
	StatusRemarks:    "case_status.status_remarks",
	UserID:           "case_status.user_id",
	CreatedAt:        "case_status.created_at",
	UpdatedAt:        "case_status.updated_at",
	DeletedAt:        "case_status.deleted_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

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

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var CaseStatusWhere = struct {
	ID               whereHelperint
	IncidentReportID whereHelperint
	Status           whereHelperstring
	StatusRemarks    whereHelperstring
	UserID           whereHelperint
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpernull_Time
	DeletedAt        whereHelpernull_Time
}{
	ID:               whereHelperint{field: "`case_status`.`id`"},
	IncidentReportID: whereHelperint{field: "`case_status`.`incident_report_id`"},
	Status:           whereHelperstring{field: "`case_status`.`status`"},
	StatusRemarks:    whereHelperstring{field: "`case_status`.`status_remarks`"},
	UserID:           whereHelperint{field: "`case_status`.`user_id`"},
	CreatedAt:        whereHelpertime_Time{field: "`case_status`.`created_at`"},
	UpdatedAt:        whereHelpernull_Time{field: "`case_status`.`updated_at`"},
	DeletedAt:        whereHelpernull_Time{field: "`case_status`.`deleted_at`"},
}

// CaseStatusRels is where relationship names are stored.
var CaseStatusRels = struct {
	IncidentReport string
	User           string
}{
	IncidentReport: "IncidentReport",
	User:           "User",
}

// caseStatusR is where relationships are stored.
type caseStatusR struct {
	IncidentReport *IncidentReport `boil:"IncidentReport" json:"IncidentReport" toml:"IncidentReport" yaml:"IncidentReport"`
	User           *Employee       `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*caseStatusR) NewStruct() *caseStatusR {
	return &caseStatusR{}
}

func (r *caseStatusR) GetIncidentReport() *IncidentReport {
	if r == nil {
		return nil
	}
	return r.IncidentReport
}

func (r *caseStatusR) GetUser() *Employee {
	if r == nil {
		return nil
	}
	return r.User
}

// caseStatusL is where Load methods for each relationship are stored.
type caseStatusL struct{}

var (
	caseStatusAllColumns            = []string{"id", "incident_report_id", "status", "status_remarks", "user_id", "created_at", "updated_at", "deleted_at"}
	caseStatusColumnsWithoutDefault = []string{"incident_report_id", "status", "status_remarks", "user_id", "updated_at", "deleted_at"}
	caseStatusColumnsWithDefault    = []string{"id", "created_at"}
	caseStatusPrimaryKeyColumns     = []string{"id"}
	caseStatusGeneratedColumns      = []string{}
)

type (
	// CaseStatusSlice is an alias for a slice of pointers to CaseStatus.
	// This should almost always be used instead of []CaseStatus.
	CaseStatusSlice []*CaseStatus

	caseStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	caseStatusType                 = reflect.TypeOf(&CaseStatus{})
	caseStatusMapping              = queries.MakeStructMapping(caseStatusType)
	caseStatusPrimaryKeyMapping, _ = queries.BindMapping(caseStatusType, caseStatusMapping, caseStatusPrimaryKeyColumns)
	caseStatusInsertCacheMut       sync.RWMutex
	caseStatusInsertCache          = make(map[string]insertCache)
	caseStatusUpdateCacheMut       sync.RWMutex
	caseStatusUpdateCache          = make(map[string]updateCache)
	caseStatusUpsertCacheMut       sync.RWMutex
	caseStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single caseStatus record from the query.
func (q caseStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CaseStatus, error) {
	o := &CaseStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for case_status")
	}

	return o, nil
}

// All returns all CaseStatus records from the query.
func (q caseStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (CaseStatusSlice, error) {
	var o []*CaseStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CaseStatus slice")
	}

	return o, nil
}

// Count returns the count of all CaseStatus records in the query.
func (q caseStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count case_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q caseStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if case_status exists")
	}

	return count > 0, nil
}

// IncidentReport pointed to by the foreign key.
func (o *CaseStatus) IncidentReport(mods ...qm.QueryMod) incidentReportQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.IncidentReportID),
	}

	queryMods = append(queryMods, mods...)

	return IncidentReports(queryMods...)
}

// User pointed to by the foreign key.
func (o *CaseStatus) User(mods ...qm.QueryMod) employeeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Employees(queryMods...)
}

// LoadIncidentReport allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (caseStatusL) LoadIncidentReport(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCaseStatus interface{}, mods queries.Applicator) error {
	var slice []*CaseStatus
	var object *CaseStatus

	if singular {
		object = maybeCaseStatus.(*CaseStatus)
	} else {
		slice = *maybeCaseStatus.(*[]*CaseStatus)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &caseStatusR{}
		}
		args = append(args, object.IncidentReportID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &caseStatusR{}
			}

			for _, a := range args {
				if a == obj.IncidentReportID {
					continue Outer
				}
			}

			args = append(args, obj.IncidentReportID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`incident_reports`),
		qm.WhereIn(`incident_reports.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load IncidentReport")
	}

	var resultSlice []*IncidentReport
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice IncidentReport")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for incident_reports")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for incident_reports")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.IncidentReport = foreign
		if foreign.R == nil {
			foreign.R = &incidentReportR{}
		}
		foreign.R.CaseStatuses = append(foreign.R.CaseStatuses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.IncidentReportID == foreign.ID {
				local.R.IncidentReport = foreign
				if foreign.R == nil {
					foreign.R = &incidentReportR{}
				}
				foreign.R.CaseStatuses = append(foreign.R.CaseStatuses, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (caseStatusL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCaseStatus interface{}, mods queries.Applicator) error {
	var slice []*CaseStatus
	var object *CaseStatus

	if singular {
		object = maybeCaseStatus.(*CaseStatus)
	} else {
		slice = *maybeCaseStatus.(*[]*CaseStatus)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &caseStatusR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &caseStatusR{}
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
		qm.From(`employees`),
		qm.WhereIn(`employees.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Employee")
	}

	var resultSlice []*Employee
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Employee")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for employees")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for employees")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &employeeR{}
		}
		foreign.R.UserCaseStatuses = append(foreign.R.UserCaseStatuses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &employeeR{}
				}
				foreign.R.UserCaseStatuses = append(foreign.R.UserCaseStatuses, local)
				break
			}
		}
	}

	return nil
}

// SetIncidentReport of the caseStatus to the related item.
// Sets o.R.IncidentReport to related.
// Adds o to related.R.CaseStatuses.
func (o *CaseStatus) SetIncidentReport(ctx context.Context, exec boil.ContextExecutor, insert bool, related *IncidentReport) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `case_status` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"incident_report_id"}),
		strmangle.WhereClause("`", "`", 0, caseStatusPrimaryKeyColumns),
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

	o.IncidentReportID = related.ID
	if o.R == nil {
		o.R = &caseStatusR{
			IncidentReport: related,
		}
	} else {
		o.R.IncidentReport = related
	}

	if related.R == nil {
		related.R = &incidentReportR{
			CaseStatuses: CaseStatusSlice{o},
		}
	} else {
		related.R.CaseStatuses = append(related.R.CaseStatuses, o)
	}

	return nil
}

// SetUser of the caseStatus to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserCaseStatuses.
func (o *CaseStatus) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Employee) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `case_status` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, caseStatusPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &caseStatusR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &employeeR{
			UserCaseStatuses: CaseStatusSlice{o},
		}
	} else {
		related.R.UserCaseStatuses = append(related.R.UserCaseStatuses, o)
	}

	return nil
}

// CaseStatuses retrieves all the records using an executor.
func CaseStatuses(mods ...qm.QueryMod) caseStatusQuery {
	mods = append(mods, qm.From("`case_status`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`case_status`.*"})
	}

	return caseStatusQuery{q}
}

// FindCaseStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCaseStatus(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CaseStatus, error) {
	caseStatusObj := &CaseStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `case_status` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, caseStatusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from case_status")
	}

	return caseStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CaseStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no case_status provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(caseStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	caseStatusInsertCacheMut.RLock()
	cache, cached := caseStatusInsertCache[key]
	caseStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			caseStatusAllColumns,
			caseStatusColumnsWithDefault,
			caseStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(caseStatusType, caseStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(caseStatusType, caseStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `case_status` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `case_status` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `case_status` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, caseStatusPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into case_status")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == caseStatusMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for case_status")
	}

CacheNoHooks:
	if !cached {
		caseStatusInsertCacheMut.Lock()
		caseStatusInsertCache[key] = cache
		caseStatusInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CaseStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CaseStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	caseStatusUpdateCacheMut.RLock()
	cache, cached := caseStatusUpdateCache[key]
	caseStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			caseStatusAllColumns,
			caseStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update case_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `case_status` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, caseStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(caseStatusType, caseStatusMapping, append(wl, caseStatusPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update case_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for case_status")
	}

	if !cached {
		caseStatusUpdateCacheMut.Lock()
		caseStatusUpdateCache[key] = cache
		caseStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q caseStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for case_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for case_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CaseStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), caseStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `case_status` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, caseStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in caseStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all caseStatus")
	}
	return rowsAff, nil
}

var mySQLCaseStatusUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CaseStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no case_status provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(caseStatusColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCaseStatusUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	caseStatusUpsertCacheMut.RLock()
	cache, cached := caseStatusUpsertCache[key]
	caseStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			caseStatusAllColumns,
			caseStatusColumnsWithDefault,
			caseStatusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			caseStatusAllColumns,
			caseStatusPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert case_status, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`case_status`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `case_status` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(caseStatusType, caseStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(caseStatusType, caseStatusMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for case_status")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == caseStatusMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(caseStatusType, caseStatusMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for case_status")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for case_status")
	}

CacheNoHooks:
	if !cached {
		caseStatusUpsertCacheMut.Lock()
		caseStatusUpsertCache[key] = cache
		caseStatusUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CaseStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CaseStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CaseStatus provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), caseStatusPrimaryKeyMapping)
	sql := "DELETE FROM `case_status` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from case_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for case_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q caseStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no caseStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from case_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for case_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CaseStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), caseStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `case_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, caseStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from caseStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for case_status")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CaseStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCaseStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CaseStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CaseStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), caseStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `case_status`.* FROM `case_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, caseStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CaseStatusSlice")
	}

	*o = slice

	return nil
}

// CaseStatusExists checks if the CaseStatus row exists.
func CaseStatusExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `case_status` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if case_status exists")
	}

	return exists, nil
}
