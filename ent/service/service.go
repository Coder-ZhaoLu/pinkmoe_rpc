// Code generated by ent, DO NOT EDIT.

package service

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the service type in the database.
	Label = "service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldAuthorUUID holds the string denoting the author_uuid field in the database.
	FieldAuthorUUID = "author_uuid"
	// FieldCover holds the string denoting the cover field in the database.
	FieldCover = "cover"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldView holds the string denoting the view field in the database.
	FieldView = "view"
	// Table holds the table name of the service in the database.
	Table = "services"
)

// Columns holds all SQL columns for service fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldTitle,
	FieldContent,
	FieldCategoryID,
	FieldAuthorUUID,
	FieldCover,
	FieldType,
	FieldPrice,
	FieldView,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType uint32
	// DefaultView holds the default value on creation for the "view" field.
	DefaultView uint64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)