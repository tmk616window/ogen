// Code generated by ent, DO NOT EDIT.

package label

import (
	"server/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Label {
	return predicate.Label(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Label {
	return predicate.Label(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Label {
	return predicate.Label(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Label {
	return predicate.Label(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Label {
	return predicate.Label(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Label {
	return predicate.Label(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Label {
	return predicate.Label(sql.FieldLTE(FieldID, id))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldValue, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldUpdatedAt, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Label {
	return predicate.Label(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Label {
	return predicate.Label(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Label {
	return predicate.Label(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Label {
	return predicate.Label(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Label {
	return predicate.Label(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Label {
	return predicate.Label(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Label {
	return predicate.Label(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Label {
	return predicate.Label(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Label {
	return predicate.Label(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Label {
	return predicate.Label(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Label {
	return predicate.Label(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Label {
	return predicate.Label(sql.FieldContainsFold(FieldValue, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Label {
	return predicate.Label(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Label {
	return predicate.Label(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Label {
	return predicate.Label(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Label {
	return predicate.Label(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Label {
	return predicate.Label(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasTodos applies the HasEdge predicate on the "todos" edge.
func HasTodos() predicate.Label {
	return predicate.Label(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TodosTable, TodosPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTodosWith applies the HasEdge predicate on the "todos" edge with a given conditions (other predicates).
func HasTodosWith(preds ...predicate.Todo) predicate.Label {
	return predicate.Label(func(s *sql.Selector) {
		step := newTodosStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Label) predicate.Label {
	return predicate.Label(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Label) predicate.Label {
	return predicate.Label(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Label) predicate.Label {
	return predicate.Label(sql.NotPredicates(p))
}
