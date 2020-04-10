// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/role"
	"github.com/confus1on/UKM/ent/schema"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/confus1on/UKM/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescFirstName is the schema descriptor for firstName field.
	profileDescFirstName := profileFields[0].Descriptor()
	// profile.FirstNameValidator is a validator for the "firstName" field. It is called by the builders before save.
	profile.FirstNameValidator = profileDescFirstName.Validators[0].(func(string) error)
	// profileDescLastName is the schema descriptor for lastName field.
	profileDescLastName := profileFields[1].Descriptor()
	// profile.LastNameValidator is a validator for the "lastName" field. It is called by the builders before save.
	profile.LastNameValidator = profileDescLastName.Validators[0].(func(string) error)
	// profileDescPhone is the schema descriptor for phone field.
	profileDescPhone := profileFields[6].Descriptor()
	// profile.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	profile.PhoneValidator = profileDescPhone.Validators[0].(func(string) error)
	// profileDescStatus is the schema descriptor for status field.
	profileDescStatus := profileFields[7].Descriptor()
	// profile.DefaultStatus holds the default value on creation for the status field.
	profile.DefaultStatus = profileDescStatus.Default.(bool)
	// profileDescImgProfile is the schema descriptor for img_profile field.
	profileDescImgProfile := profileFields[8].Descriptor()
	// profile.DefaultImgProfile holds the default value on creation for the img_profile field.
	profile.DefaultImgProfile = profileDescImgProfile.Default.(string)
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileFields[9].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescUpdatedAt is the schema descriptor for updated_at field.
	profileDescUpdatedAt := profileFields[10].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	// profile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	profile.UpdateDefaultUpdatedAt = profileDescUpdatedAt.UpdateDefault.(func() time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescValue is the schema descriptor for value field.
	roleDescValue := roleFields[0].Descriptor()
	// role.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	role.ValueValidator = roleDescValue.Validators[0].(func(string) error)
	ukmFields := schema.Ukm{}.Fields()
	_ = ukmFields
	// ukmDescName is the schema descriptor for name field.
	ukmDescName := ukmFields[0].Descriptor()
	// ukm.NameValidator is a validator for the "name" field. It is called by the builders before save.
	ukm.NameValidator = ukmDescName.Validators[0].(func(string) error)
	// ukmDescCreatedAt is the schema descriptor for created_at field.
	ukmDescCreatedAt := ukmFields[1].Descriptor()
	// ukm.DefaultCreatedAt holds the default value on creation for the created_at field.
	ukm.DefaultCreatedAt = ukmDescCreatedAt.Default.(func() time.Time)
	// ukmDescUpdatedAt is the schema descriptor for updated_at field.
	ukmDescUpdatedAt := ukmFields[2].Descriptor()
	// ukm.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ukm.DefaultUpdatedAt = ukmDescUpdatedAt.Default.(func() time.Time)
	// ukm.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ukm.UpdateDefaultUpdatedAt = ukmDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
