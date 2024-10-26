package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
		field.String("name").NotEmpty(),
		field.Time("finished_at").Nillable(),
		field.Int("priority_id"),
		field.Int("status_id").Default(1),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("priority", Priority.Type).
			Ref("todo").
			Required().
			Unique().
			Field("priority_id"),
		edge.From("status", Status.Type).
			Ref("todo").
			Required().
			Unique().
			Field("status_id"),
	}
}
