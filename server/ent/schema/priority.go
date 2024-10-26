package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Priority holds the schema definition for the Priority entity.
type Priority struct {
	ent.Schema
}

// Fields of the Priority.
func (Priority) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now()),
	}
}

// Edges of the Priority.
func (Priority) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todo", Todo.Type),
	}
}
