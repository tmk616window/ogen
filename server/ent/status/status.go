// Code generated by ent, DO NOT EDIT.

package status

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the status type in the database.
	Label = "status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"
	// Table holds the table name of the status in the database.
	Table = "status"
	// TodoTable is the table that holds the todo relation/edge.
	TodoTable = "todos"
	// TodoInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodoInverseTable = "todos"
	// TodoColumn is the table column denoting the todo relation/edge.
	TodoColumn = "status_id"
)

// Columns holds all SQL columns for status fields.
var Columns = []string{
	FieldID,
	FieldValue,
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
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
)

// OrderOption defines the ordering options for the Status queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByTodoField orders the results by todo field.
func ByTodoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodoStep(), sql.OrderByField(field, opts...))
	}
}
func newTodoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, TodoTable, TodoColumn),
	)
}