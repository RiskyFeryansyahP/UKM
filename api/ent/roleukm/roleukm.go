// Code generated by entc, DO NOT EDIT.

package roleukm

const (
	// Label holds the string label denoting the roleukm type in the database.
	Label = "role_ukm"
	// FieldID holds the string denoting the id field in the database.
	FieldID         = "id" // FieldStatusRole holds the string denoting the status_role vertex property in the database.
	FieldStatusRole = "status_role"

	// EdgeProfileRoles holds the string denoting the profile_roles edge name in mutations.
	EdgeProfileRoles = "profile_roles"

	// Table holds the table name of the roleukm in the database.
	Table = "role_ukm"
	// ProfileRolesTable is the table the holds the profile_roles relation/edge.
	ProfileRolesTable = "profile_ukm"
	// ProfileRolesInverseTable is the table name for the ProfileUKM entity.
	// It exists in this package in order to avoid circular dependency with the "profileukm" package.
	ProfileRolesInverseTable = "profile_ukm"
	// ProfileRolesColumn is the table column denoting the profile_roles relation/edge.
	ProfileRolesColumn = "role_ukm_profile_roles"
)

// Columns holds all SQL columns for roleukm fields.
var Columns = []string{
	FieldID,
	FieldStatusRole,
}

var (
	// StatusRoleValidator is a validator for the "status_role" field. It is called by the builders before save.
	StatusRoleValidator func(string) error
)