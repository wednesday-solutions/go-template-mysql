// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testIncidentReportAttachments(t *testing.T) {
	t.Parallel()

	query := IncidentReportAttachments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIncidentReportAttachmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
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

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIncidentReportAttachmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := IncidentReportAttachments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIncidentReportAttachmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IncidentReportAttachmentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIncidentReportAttachmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IncidentReportAttachmentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if IncidentReportAttachment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IncidentReportAttachmentExists to return true, but got false.")
	}
}

func testIncidentReportAttachmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	incidentReportAttachmentFound, err := FindIncidentReportAttachment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if incidentReportAttachmentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIncidentReportAttachmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = IncidentReportAttachments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIncidentReportAttachmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := IncidentReportAttachments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIncidentReportAttachmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	incidentReportAttachmentOne := &IncidentReportAttachment{}
	incidentReportAttachmentTwo := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, incidentReportAttachmentOne, incidentReportAttachmentDBTypes, false, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}
	if err = randomize.Struct(seed, incidentReportAttachmentTwo, incidentReportAttachmentDBTypes, false, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = incidentReportAttachmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = incidentReportAttachmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IncidentReportAttachments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIncidentReportAttachmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	incidentReportAttachmentOne := &IncidentReportAttachment{}
	incidentReportAttachmentTwo := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, incidentReportAttachmentOne, incidentReportAttachmentDBTypes, false, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}
	if err = randomize.Struct(seed, incidentReportAttachmentTwo, incidentReportAttachmentDBTypes, false, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = incidentReportAttachmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = incidentReportAttachmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testIncidentReportAttachmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIncidentReportAttachmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(incidentReportAttachmentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIncidentReportAttachmentToOneIncidentReportUsingIncidentReport(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local IncidentReportAttachment
	var foreign IncidentReport

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, incidentReportAttachmentDBTypes, false, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, incidentReportDBTypes, false, incidentReportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReport struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.IncidentReportID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.IncidentReport().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := IncidentReportAttachmentSlice{&local}
	if err = local.L.LoadIncidentReport(ctx, tx, false, (*[]*IncidentReportAttachment)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.IncidentReport == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.IncidentReport = nil
	if err = local.L.LoadIncidentReport(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.IncidentReport == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIncidentReportAttachmentToOneSetOpIncidentReportUsingIncidentReport(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IncidentReportAttachment
	var b, c IncidentReport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, incidentReportAttachmentDBTypes, false, strmangle.SetComplement(incidentReportAttachmentPrimaryKeyColumns, incidentReportAttachmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, incidentReportDBTypes, false, strmangle.SetComplement(incidentReportPrimaryKeyColumns, incidentReportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, incidentReportDBTypes, false, strmangle.SetComplement(incidentReportPrimaryKeyColumns, incidentReportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*IncidentReport{&b, &c} {
		err = a.SetIncidentReport(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.IncidentReport != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.IncidentReportAttachments[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.IncidentReportID != x.ID {
			t.Error("foreign key was wrong value", a.IncidentReportID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.IncidentReportID))
		reflect.Indirect(reflect.ValueOf(&a.IncidentReportID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.IncidentReportID != x.ID {
			t.Error("foreign key was wrong value", a.IncidentReportID, x.ID)
		}
	}
}

func testIncidentReportAttachmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
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

func testIncidentReportAttachmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IncidentReportAttachmentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIncidentReportAttachmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IncidentReportAttachments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	incidentReportAttachmentDBTypes = map[string]string{`ID`: `int`, `IncidentReportID`: `int`, `URL`: `text`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                               = bytes.MinRead
)

func testIncidentReportAttachmentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(incidentReportAttachmentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(incidentReportAttachmentAllColumns) == len(incidentReportAttachmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIncidentReportAttachmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(incidentReportAttachmentAllColumns) == len(incidentReportAttachmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IncidentReportAttachment{}
	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, incidentReportAttachmentDBTypes, true, incidentReportAttachmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(incidentReportAttachmentAllColumns, incidentReportAttachmentPrimaryKeyColumns) {
		fields = incidentReportAttachmentAllColumns
	} else {
		fields = strmangle.SetComplement(
			incidentReportAttachmentAllColumns,
			incidentReportAttachmentPrimaryKeyColumns,
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

	slice := IncidentReportAttachmentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIncidentReportAttachmentsUpsert(t *testing.T) {
	t.Parallel()

	if len(incidentReportAttachmentAllColumns) == len(incidentReportAttachmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLIncidentReportAttachmentUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := IncidentReportAttachment{}
	if err = randomize.Struct(seed, &o, incidentReportAttachmentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IncidentReportAttachment: %s", err)
	}

	count, err := IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, incidentReportAttachmentDBTypes, false, incidentReportAttachmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IncidentReportAttachment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IncidentReportAttachment: %s", err)
	}

	count, err = IncidentReportAttachments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}