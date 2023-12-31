// Code generated by ent, DO NOT EDIT.

package keys

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the keys type in the database.
	Label = "keys"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVerificationKeys holds the string denoting the verification_keys field in the database.
	FieldVerificationKeys = "verification_keys"
	// FieldSigningKey holds the string denoting the signing_key field in the database.
	FieldSigningKey = "signing_key"
	// FieldSigningKeyPub holds the string denoting the signing_key_pub field in the database.
	FieldSigningKeyPub = "signing_key_pub"
	// FieldNextRotation holds the string denoting the next_rotation field in the database.
	FieldNextRotation = "next_rotation"
	// Table holds the table name of the keys in the database.
	Table = "keys"
)

// Columns holds all SQL columns for keys fields.
var Columns = []string{
	FieldID,
	FieldVerificationKeys,
	FieldSigningKey,
	FieldSigningKeyPub,
	FieldNextRotation,
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Keys queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNextRotation orders the results by the next_rotation field.
func ByNextRotation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextRotation, opts...).ToFunc()
}
