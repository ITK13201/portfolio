// Code generated by entc, DO NOT EDIT.

package work

import (
	"time"
)

const (
	// Label holds the string label denoting the work type in the database.
	Label = "work"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescriptionJp holds the string denoting the description_jp field in the database.
	FieldDescriptionJp = "description_jp"
	// FieldDescriptionEn holds the string denoting the description_en field in the database.
	FieldDescriptionEn = "description_en"
	// FieldLanguageID holds the string denoting the language_id field in the database.
	FieldLanguageID = "language_id"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeLanguage holds the string denoting the language edge name in mutations.
	EdgeLanguage = "language"
	// Table holds the table name of the work in the database.
	Table = "works"
	// LanguageTable is the table that holds the language relation/edge.
	LanguageTable = "works"
	// LanguageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	LanguageInverseTable = "images"
	// LanguageColumn is the table column denoting the language relation/edge.
	LanguageColumn = "language_id"
)

// Columns holds all SQL columns for work fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescriptionJp,
	FieldDescriptionEn,
	FieldLanguageID,
	FieldLink,
	FieldPriority,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionJpValidator is a validator for the "description_jp" field. It is called by the builders before save.
	DescriptionJpValidator func(string) error
	// DescriptionEnValidator is a validator for the "description_en" field. It is called by the builders before save.
	DescriptionEnValidator func(string) error
	// LinkValidator is a validator for the "link" field. It is called by the builders before save.
	LinkValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
