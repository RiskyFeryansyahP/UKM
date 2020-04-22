// Code generated by entc, DO NOT EDIT.

package ukm

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the ukm type in the database.
	Label = "ukm"
	// FieldID holds the string denoting the id field in the database.
	FieldID        = "id"         // FieldName holds the string denoting the name vertex property in the database.
	FieldName      = "name"       // FieldStatus holds the string denoting the status vertex property in the database.
	FieldStatus    = "status"     // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at" // FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeProfiles holds the string denoting the profiles edge name in mutations.
	EdgeProfiles = "profiles"

	// Table holds the table name of the ukm in the database.
	Table = "ukms"
	// ProfilesTable is the table the holds the profiles relation/edge. The primary key declared below.
	ProfilesTable = "profile_ukm"
	// ProfilesInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfilesInverseTable = "profiles"
)

// Columns holds all SQL columns for ukm fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// ProfilesPrimaryKey and ProfilesColumn2 are the table columns denoting the
	// primary key for the profiles relation (M2M).
	ProfilesPrimaryKey = []string{"profile_id", "ukm_id"}
)

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusOpen  Status = "open"
	StatusClose Status = "close"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOpen, StatusClose:
		return nil
	default:
		return fmt.Errorf("ukm: invalid enum value for status field: %q", s)
	}
}
