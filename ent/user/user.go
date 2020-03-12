// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/fitzix/assassin/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password vertex property in the database.
	FieldPassword = "password"
	// FieldCode holds the string denoting the code vertex property in the database.
	FieldCode = "code"
	// FieldStatus holds the string denoting the status vertex property in the database.
	FieldStatus = "status"

	// Table holds the table name of the user in the database.
	Table = "users"
	// RoleTable is the table the holds the role relation/edge.
	RoleTable = "users"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "user_role"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPassword,
	FieldCode,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"user_role",
}

var (
	fields = schema.User{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[0].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = func() func(string) error {
		validators := descName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descPassword is the schema descriptor for password field.
	descPassword = fields[1].Descriptor()
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator = descPassword.Validators[0].(func(string) error)

	// descCode is the schema descriptor for code field.
	descCode = fields[2].Descriptor()
	// DefaultCode holds the default value on creation for the code field.
	DefaultCode = descCode.Default.(uint)
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator = descCode.Validators[0].(func(uint) error)
)

// Status defines the type for the status enum field.
type Status string

// StatusNormal is the default Status.
const DefaultStatus = StatusNormal

// Status values.
const (
	StatusNormal   Status = "normal"
	StatusAbnormal Status = "abnormal"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusNormal, StatusAbnormal:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}
