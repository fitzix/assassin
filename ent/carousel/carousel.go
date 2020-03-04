// Code generated by entc, DO NOT EDIT.

package carousel

import (
	"github.com/fitzix/assassin/ent/schema"
)

const (
	// Label holds the string label denoting the carousel type in the database.
	Label = "carousel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url vertex property in the database.
	FieldURL = "url"

	// Table holds the table name of the carousel in the database.
	Table = "carousels"
)

// Columns holds all SQL columns for carousel fields.
var Columns = []string{
	FieldID,
	FieldURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Carousel type.
var ForeignKeys = []string{
	"app_carousels",
}

var (
	fields = schema.Carousel{}.Fields()

	// descURL is the schema descriptor for url field.
	descURL = fields[0].Descriptor()
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator = descURL.Validators[0].(func(string) error)
)