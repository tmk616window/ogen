// Code generated by ent, DO NOT EDIT.

package todo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldPriorityID holds the string denoting the priority_id field in the database.
	FieldPriorityID = "priority_id"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePriority holds the string denoting the priority edge name in mutations.
	EdgePriority = "priority"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// PriorityTable is the table that holds the priority relation/edge.
	PriorityTable = "todos"
	// PriorityInverseTable is the table name for the Priority entity.
	// It exists in this package in order to avoid circular dependency with the "priority" package.
	PriorityInverseTable = "priorities"
	// PriorityColumn is the table column denoting the priority relation/edge.
	PriorityColumn = "priority_id"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "todos"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "status_id"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldName,
	FieldFinishedAt,
	FieldPriorityID,
	FieldStatusID,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultStatusID holds the default value on creation for the "status_id" field.
	DefaultStatusID int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Todo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByPriorityID orders the results by the priority_id field.
func ByPriorityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriorityID, opts...).ToFunc()
}

// ByStatusID orders the results by the status_id field.
func ByStatusID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPriorityField orders the results by priority field.
func ByPriorityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPriorityStep(), sql.OrderByField(field, opts...))
	}
}

// ByStatusField orders the results by status field.
func ByStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusStep(), sql.OrderByField(field, opts...))
	}
}
func newPriorityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PriorityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PriorityTable, PriorityColumn),
	)
}
func newStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StatusTable, StatusColumn),
	)
}
