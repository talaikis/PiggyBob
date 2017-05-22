// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testExpenseCategories(t *testing.T) {
	t.Parallel()

	query := ExpenseCategories(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testExpenseCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = expenseCategory.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExpenseCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ExpenseCategories(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExpenseCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ExpenseCategorySlice{expenseCategory}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testExpenseCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ExpenseCategoryExists(tx, expenseCategory.ID)
	if err != nil {
		t.Errorf("Unable to check if ExpenseCategory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ExpenseCategoryExistsG to return true, but got false.")
	}
}
func testExpenseCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	expenseCategoryFound, err := FindExpenseCategory(tx, expenseCategory.ID)
	if err != nil {
		t.Error(err)
	}

	if expenseCategoryFound == nil {
		t.Error("want a record, got nil")
	}
}
func testExpenseCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ExpenseCategories(tx).Bind(expenseCategory); err != nil {
		t.Error(err)
	}
}

func testExpenseCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := ExpenseCategories(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testExpenseCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategoryOne := &ExpenseCategory{}
	expenseCategoryTwo := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategoryOne, expenseCategoryDBTypes, false, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, expenseCategoryTwo, expenseCategoryDBTypes, false, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = expenseCategoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ExpenseCategories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testExpenseCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	expenseCategoryOne := &ExpenseCategory{}
	expenseCategoryTwo := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategoryOne, expenseCategoryDBTypes, false, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, expenseCategoryTwo, expenseCategoryDBTypes, false, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = expenseCategoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func expenseCategoryBeforeInsertHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryAfterInsertHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryAfterSelectHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryBeforeUpdateHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryAfterUpdateHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryBeforeDeleteHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryAfterDeleteHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryBeforeUpsertHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func expenseCategoryAfterUpsertHook(e boil.Executor, o *ExpenseCategory) error {
	*o = ExpenseCategory{}
	return nil
}

func testExpenseCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &ExpenseCategory{}
	o := &ExpenseCategory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, expenseCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory object: %s", err)
	}

	AddExpenseCategoryHook(boil.BeforeInsertHook, expenseCategoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	expenseCategoryBeforeInsertHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.AfterInsertHook, expenseCategoryAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	expenseCategoryAfterInsertHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.AfterSelectHook, expenseCategoryAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	expenseCategoryAfterSelectHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.BeforeUpdateHook, expenseCategoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	expenseCategoryBeforeUpdateHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.AfterUpdateHook, expenseCategoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	expenseCategoryAfterUpdateHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.BeforeDeleteHook, expenseCategoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	expenseCategoryBeforeDeleteHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.AfterDeleteHook, expenseCategoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	expenseCategoryAfterDeleteHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.BeforeUpsertHook, expenseCategoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	expenseCategoryBeforeUpsertHooks = []ExpenseCategoryHook{}

	AddExpenseCategoryHook(boil.AfterUpsertHook, expenseCategoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	expenseCategoryAfterUpsertHooks = []ExpenseCategoryHook{}
}
func testExpenseCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExpenseCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx, expenseCategoryColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExpenseCategoryToManyCategoryExpenses(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ExpenseCategory
	var b, c Expense

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, expenseDBTypes, false, expenseColumnsWithDefault...)
	randomize.Struct(seed, &c, expenseDBTypes, false, expenseColumnsWithDefault...)

	b.CategoryID = a.ID
	c.CategoryID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	expense, err := a.CategoryExpenses(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range expense {
		if v.CategoryID == b.CategoryID {
			bFound = true
		}
		if v.CategoryID == c.CategoryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ExpenseCategorySlice{&a}
	if err = a.L.LoadCategoryExpenses(tx, false, (*[]*ExpenseCategory)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryExpenses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CategoryExpenses = nil
	if err = a.L.LoadCategoryExpenses(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CategoryExpenses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", expense)
	}
}

func testExpenseCategoryToManyAddOpCategoryExpenses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ExpenseCategory
	var b, c, d, e Expense

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, expenseCategoryDBTypes, false, strmangle.SetComplement(expenseCategoryPrimaryKeyColumns, expenseCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Expense{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, expenseDBTypes, false, strmangle.SetComplement(expensePrimaryKeyColumns, expenseColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Expense{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCategoryExpenses(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CategoryID {
			t.Error("foreign key was wrong value", a.ID, first.CategoryID)
		}
		if a.ID != second.CategoryID {
			t.Error("foreign key was wrong value", a.ID, second.CategoryID)
		}

		if first.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Category != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CategoryExpenses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CategoryExpenses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CategoryExpenses(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testExpenseCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = expenseCategory.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testExpenseCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ExpenseCategorySlice{expenseCategory}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testExpenseCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ExpenseCategories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	expenseCategoryDBTypes = map[string]string{`Description`: `character varying`, `ID`: `integer`, `Title`: `character varying`}
	_                      = bytes.MinRead
)

func testExpenseCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if len(expenseCategoryColumns) == len(expenseCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	if err = expenseCategory.Update(tx); err != nil {
		t.Error(err)
	}
}

func testExpenseCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(expenseCategoryColumns) == len(expenseCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	expenseCategory := &ExpenseCategory{}
	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, expenseCategory, expenseCategoryDBTypes, true, expenseCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(expenseCategoryColumns, expenseCategoryPrimaryKeyColumns) {
		fields = expenseCategoryColumns
	} else {
		fields = strmangle.SetComplement(
			expenseCategoryColumns,
			expenseCategoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(expenseCategory))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ExpenseCategorySlice{expenseCategory}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testExpenseCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(expenseCategoryColumns) == len(expenseCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	expenseCategory := ExpenseCategory{}
	if err = randomize.Struct(seed, &expenseCategory, expenseCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = expenseCategory.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert ExpenseCategory: %s", err)
	}

	count, err := ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &expenseCategory, expenseCategoryDBTypes, false, expenseCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExpenseCategory struct: %s", err)
	}

	if err = expenseCategory.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert ExpenseCategory: %s", err)
	}

	count, err = ExpenseCategories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}