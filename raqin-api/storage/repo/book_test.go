// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repo

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

func testBooks(t *testing.T) {
	t.Parallel()

	query := Books()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBooksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
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

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBooksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Books().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBooksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBooksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Book exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookExists to return true, but got false.")
	}
}

func testBooksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookFound, err := FindBook(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBooksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Books().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBooksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Books().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBooksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookOne := &Book{}
	bookTwo := &Book{}
	if err = randomize.Struct(seed, bookOne, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}
	if err = randomize.Struct(seed, bookTwo, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Books().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBooksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookOne := &Book{}
	bookTwo := &Book{}
	if err = randomize.Struct(seed, bookOne, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}
	if err = randomize.Struct(seed, bookTwo, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func bookAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Book) error {
	*o = Book{}
	return nil
}

func testBooksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Book{}
	o := &Book{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Book object: %s", err)
	}

	AddBookHook(boil.BeforeInsertHook, bookBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookBeforeInsertHooks = []BookHook{}

	AddBookHook(boil.AfterInsertHook, bookAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookAfterInsertHooks = []BookHook{}

	AddBookHook(boil.AfterSelectHook, bookAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookAfterSelectHooks = []BookHook{}

	AddBookHook(boil.BeforeUpdateHook, bookBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookBeforeUpdateHooks = []BookHook{}

	AddBookHook(boil.AfterUpdateHook, bookAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookAfterUpdateHooks = []BookHook{}

	AddBookHook(boil.BeforeDeleteHook, bookBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookBeforeDeleteHooks = []BookHook{}

	AddBookHook(boil.AfterDeleteHook, bookAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookAfterDeleteHooks = []BookHook{}

	AddBookHook(boil.BeforeUpsertHook, bookBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookBeforeUpsertHooks = []BookHook{}

	AddBookHook(boil.AfterUpsertHook, bookAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookAfterUpsertHooks = []BookHook{}
}

func testBooksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBooksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookToManyBookAuthors(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c BookAuthor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookAuthorDBTypes, false, bookAuthorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookAuthorDBTypes, false, bookAuthorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BookID = a.ID
	c.BookID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BookAuthors().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BookID == b.BookID {
			bFound = true
		}
		if v.BookID == c.BookID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookSlice{&a}
	if err = a.L.LoadBookAuthors(ctx, tx, false, (*[]*Book)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookAuthors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BookAuthors = nil
	if err = a.L.LoadBookAuthors(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookAuthors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookToManyBookCategories(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BookID = a.ID
	c.BookID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BookCategories().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BookID == b.BookID {
			bFound = true
		}
		if v.BookID == c.BookID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookSlice{&a}
	if err = a.L.LoadBookCategories(ctx, tx, false, (*[]*Book)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BookCategories = nil
	if err = a.L.LoadBookCategories(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookCategories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookToManyBookInitiaters(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c BookInitiater

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookInitiaterDBTypes, false, bookInitiaterColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookInitiaterDBTypes, false, bookInitiaterColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BookID = a.ID
	c.BookID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BookInitiaters().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BookID == b.BookID {
			bFound = true
		}
		if v.BookID == c.BookID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookSlice{&a}
	if err = a.L.LoadBookInitiaters(ctx, tx, false, (*[]*Book)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookInitiaters); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BookInitiaters = nil
	if err = a.L.LoadBookInitiaters(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BookInitiaters); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookToManyPages(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, pageDBTypes, false, pageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pageDBTypes, false, pageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BookID = a.ID
	c.BookID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Pages().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BookID == b.BookID {
			bFound = true
		}
		if v.BookID == c.BookID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookSlice{&a}
	if err = a.L.LoadPages(ctx, tx, false, (*[]*Book)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Pages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Pages = nil
	if err = a.L.LoadPages(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Pages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookToManyAddOpBookAuthors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c, d, e BookAuthor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookAuthor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookAuthorDBTypes, false, strmangle.SetComplement(bookAuthorPrimaryKeyColumns, bookAuthorColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BookAuthor{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBookAuthors(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BookID {
			t.Error("foreign key was wrong value", a.ID, first.BookID)
		}
		if a.ID != second.BookID {
			t.Error("foreign key was wrong value", a.ID, second.BookID)
		}

		if first.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BookAuthors[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BookAuthors[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BookAuthors().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBookToManyAddOpBookCategories(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c, d, e BookCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookCategory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BookCategory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBookCategories(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BookID {
			t.Error("foreign key was wrong value", a.ID, first.BookID)
		}
		if a.ID != second.BookID {
			t.Error("foreign key was wrong value", a.ID, second.BookID)
		}

		if first.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BookCategories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BookCategories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BookCategories().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBookToManyAddOpBookInitiaters(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c, d, e BookInitiater

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BookInitiater{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookInitiaterDBTypes, false, strmangle.SetComplement(bookInitiaterPrimaryKeyColumns, bookInitiaterColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BookInitiater{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBookInitiaters(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BookID {
			t.Error("foreign key was wrong value", a.ID, first.BookID)
		}
		if a.ID != second.BookID {
			t.Error("foreign key was wrong value", a.ID, second.BookID)
		}

		if first.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BookInitiaters[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BookInitiaters[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BookInitiaters().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBookToManyAddOpPages(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Book
	var b, c, d, e Page

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Page{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, pageDBTypes, false, strmangle.SetComplement(pagePrimaryKeyColumns, pageColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Page{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPages(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BookID {
			t.Error("foreign key was wrong value", a.ID, first.BookID)
		}
		if a.ID != second.BookID {
			t.Error("foreign key was wrong value", a.ID, second.BookID)
		}

		if first.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Book != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Pages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Pages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Pages().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBooksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
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

func testBooksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBooksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Books().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `Path`: `character varying`, `Note`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testBooksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookAllColumns) == len(bookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookDBTypes, true, bookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBooksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookAllColumns) == len(bookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Book{}
	if err = randomize.Struct(seed, o, bookDBTypes, true, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookDBTypes, true, bookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookAllColumns, bookPrimaryKeyColumns) {
		fields = bookAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookAllColumns,
			bookPrimaryKeyColumns,
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

	slice := BookSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBooksUpsert(t *testing.T) {
	t.Parallel()

	if len(bookAllColumns) == len(bookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Book{}
	if err = randomize.Struct(seed, &o, bookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Book: %s", err)
	}

	count, err := Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookDBTypes, false, bookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Book: %s", err)
	}

	count, err = Books().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
