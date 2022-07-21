// Code generated by ent, DO NOT EDIT.

package ref

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ref type in the database.
	Label = "ref"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldImmutable holds the string denoting the immutable field in the database.
	FieldImmutable = "immutable"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeRevision holds the string denoting the revision edge name in mutations.
	EdgeRevision = "revision"
	// EdgeRoutes holds the string denoting the routes edge name in mutations.
	EdgeRoutes = "routes"
	// Table holds the table name of the ref in the database.
	Table = "refs"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "refs"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_refs"
	// RevisionTable is the table that holds the revision relation/edge.
	RevisionTable = "refs"
	// RevisionInverseTable is the table name for the Revision entity.
	// It exists in this package in order to avoid circular dependency with the "revision" package.
	RevisionInverseTable = "revisions"
	// RevisionColumn is the table column denoting the revision relation/edge.
	RevisionColumn = "revision_refs"
	// RoutesTable is the table that holds the routes relation/edge.
	RoutesTable = "routes"
	// RoutesInverseTable is the table name for the Route entity.
	// It exists in this package in order to avoid circular dependency with the "route" package.
	RoutesInverseTable = "routes"
	// RoutesColumn is the table column denoting the routes relation/edge.
	RoutesColumn = "ref_routes"
)

// Columns holds all SQL columns for ref fields.
var Columns = []string{
	FieldID,
	FieldImmutable,
	FieldName,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "refs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"revision_refs",
	"workflow_refs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultImmutable holds the default value on creation for the "immutable" field.
	DefaultImmutable bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
