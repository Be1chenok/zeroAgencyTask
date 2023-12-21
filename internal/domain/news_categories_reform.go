// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package domain

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type newsCategoriesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *newsCategoriesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("news_categories").
func (v *newsCategoriesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *newsCategoriesTableType) Columns() []string {
	return []string{
		"id",
		"news_id",
		"category_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *newsCategoriesTableType) NewStruct() reform.Struct {
	return new(NewsCategories)
}

// NewRecord makes a new record for that table.
func (v *newsCategoriesTableType) NewRecord() reform.Record {
	return new(NewsCategories)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *newsCategoriesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// NewsCategoriesTable represents news_categories view or table in SQL database.
var NewsCategoriesTable = &newsCategoriesTableType{
	s: parse.StructInfo{
		Type:    "NewsCategories",
		SQLName: "news_categories",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "int", Column: "id"},
			{Name: "NewsId", Type: "int", Column: "news_id"},
			{Name: "CategoryId", Type: "int", Column: "category_id"},
		},
		PKFieldIndex: 0,
	},
	z: new(NewsCategories).Values(),
}

// String returns a string representation of this struct or record.
func (s NewsCategories) String() string {
	res := make([]string, 3)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "NewsId: " + reform.Inspect(s.NewsId, true)
	res[2] = "CategoryId: " + reform.Inspect(s.CategoryId, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *NewsCategories) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.NewsId,
		s.CategoryId,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *NewsCategories) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.NewsId,
		&s.CategoryId,
	}
}

// View returns View object for that struct.
func (s *NewsCategories) View() reform.View {
	return NewsCategoriesTable
}

// Table returns Table object for that record.
func (s *NewsCategories) Table() reform.Table {
	return NewsCategoriesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *NewsCategories) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *NewsCategories) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *NewsCategories) HasPK() bool {
	return s.Id != NewsCategoriesTable.z[NewsCategoriesTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *NewsCategories) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = NewsCategoriesTable
	_ reform.Struct = (*NewsCategories)(nil)
	_ reform.Table  = NewsCategoriesTable
	_ reform.Record = (*NewsCategories)(nil)
	_ fmt.Stringer  = (*NewsCategories)(nil)
)

func init() {
	parse.AssertUpToDate(&NewsCategoriesTable.s, new(NewsCategories))
}
