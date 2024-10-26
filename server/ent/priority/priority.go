// Code generated by ent, DO NOT EDIT.

package priority

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the priority type in the database.
	Label = "priority"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"
	// Table holds the table name of the priority in the database.
	Table = "priorities"
	// TodoTable is the table that holds the todo relation/edge.
	TodoTable = "todos"
	// TodoInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodoInverseTable = "todos"
	// TodoColumn is the table column denoting the todo relation/edge.
	TodoColumn = "priority_id"
)

// Columns holds all SQL columns for priority fields.
var Columns = []string{
	FieldID,
	FieldName,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Priority queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTodoCount orders the results by todo count.
func ByTodoCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodoStep(), opts...)
	}
}

// ByTodo orders the results by todo terms.
func ByTodo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodoStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TodoTable, TodoColumn),
	)
}
