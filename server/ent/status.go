// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"server/ent/status"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Status is the model entity for the Status schema.
type Status struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusQuery when eager-loading is set.
	Edges        StatusEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StatusEdges holds the relations/edges for other nodes in the graph.
type StatusEdges struct {
	// Todo holds the value of the todo edge.
	Todo []*Todo `json:"todo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodoOrErr returns the Todo value or an error if the edge
// was not loaded in eager-loading.
func (e StatusEdges) TodoOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todo, nil
	}
	return nil, &NotLoadedError{edge: "todo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Status) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case status.FieldID:
			values[i] = new(sql.NullInt64)
		case status.FieldValue:
			values[i] = new(sql.NullString)
		case status.FieldCreatedAt, status.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Status fields.
func (s *Status) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case status.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case status.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = value.String
			}
		case status.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case status.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Status.
// This includes values selected through modifiers, order, etc.
func (s *Status) GetValue(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTodo queries the "todo" edge of the Status entity.
func (s *Status) QueryTodo() *TodoQuery {
	return NewStatusClient(s.config).QueryTodo(s)
}

// Update returns a builder for updating this Status.
// Note that you need to call Status.Unwrap() before calling this method if this Status
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Status) Update() *StatusUpdateOne {
	return NewStatusClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Status entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Status) Unwrap() *Status {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Status is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Status) String() string {
	var builder strings.Builder
	builder.WriteString("Status(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("value=")
	builder.WriteString(s.Value)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// StatusSlice is a parsable slice of Status.
type StatusSlice []*Status
