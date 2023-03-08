// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldServiceUUID holds the string denoting the service_uuid field in the database.
	FieldServiceUUID = "service_uuid"
	// FieldUserUUID holds the string denoting the user_uuid field in the database.
	FieldUserUUID = "user_uuid"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// Table holds the table name of the comment in the database.
	Table = "comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldContent,
	FieldServiceUUID,
	FieldUserUUID,
	FieldType,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
