// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "jk", Type: field.TypeEnum, Enums: []string{"Male", "Female", "None"}},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "birth_date", Type: field.TypeString, Nullable: true},
		{Name: "year_generation", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "status", Type: field.TypeBool, Default: true},
		{Name: "img_profile", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:        "profiles",
		Columns:     ProfilesColumns,
		PrimaryKey:  []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProfileUkmColumns holds the columns for the "profile_ukm" table.
	ProfileUkmColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reason", Type: field.TypeString},
		{Name: "profile_ukms", Type: field.TypeInt, Nullable: true},
		{Name: "role_ukm_profile_roles", Type: field.TypeInt, Nullable: true},
		{Name: "ukm_profiles", Type: field.TypeInt, Nullable: true},
	}
	// ProfileUkmTable holds the schema information for the "profile_ukm" table.
	ProfileUkmTable = &schema.Table{
		Name:       "profile_ukm",
		Columns:    ProfileUkmColumns,
		PrimaryKey: []*schema.Column{ProfileUkmColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "profile_ukm_profiles_ukms",
				Columns: []*schema.Column{ProfileUkmColumns[2]},

				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "profile_ukm_role_ukm_profile_roles",
				Columns: []*schema.Column{ProfileUkmColumns[3]},

				RefColumns: []*schema.Column{RoleUkmColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "profile_ukm_ukms_profiles",
				Columns: []*schema.Column{ProfileUkmColumns[4]},

				RefColumns: []*schema.Column{UkmsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RoleUkmColumns holds the columns for the "role_ukm" table.
	RoleUkmColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status_role", Type: field.TypeString},
	}
	// RoleUkmTable holds the schema information for the "role_ukm" table.
	RoleUkmTable = &schema.Table{
		Name:        "role_ukm",
		Columns:     RoleUkmColumns,
		PrimaryKey:  []*schema.Column{RoleUkmColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UkmsColumns holds the columns for the "ukms" table.
	UkmsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"open", "close"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UkmsTable holds the schema information for the "ukms" table.
	UkmsTable = &schema.Table{
		Name:        "ukms",
		Columns:     UkmsColumns,
		PrimaryKey:  []*schema.Column{UkmsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_profile", Type: field.TypeInt, Nullable: true},
		{Name: "user_role", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_profiles_profile",
				Columns: []*schema.Column{UsersColumns[5]},

				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_roles_role",
				Columns: []*schema.Column{UsersColumns[6]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProfilesTable,
		ProfileUkmTable,
		RolesTable,
		RoleUkmTable,
		UkmsTable,
		UsersTable,
	}
)

func init() {
	ProfileUkmTable.ForeignKeys[0].RefTable = ProfilesTable
	ProfileUkmTable.ForeignKeys[1].RefTable = RoleUkmTable
	ProfileUkmTable.ForeignKeys[2].RefTable = UkmsTable
	UsersTable.ForeignKeys[0].RefTable = ProfilesTable
	UsersTable.ForeignKeys[1].RefTable = RolesTable
}
