// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/predicate"
	"server/ent/priority"
	"server/ent/todo"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PriorityUpdate is the builder for updating Priority entities.
type PriorityUpdate struct {
	config
	hooks    []Hook
	mutation *PriorityMutation
}

// Where appends a list predicates to the PriorityUpdate builder.
func (pu *PriorityUpdate) Where(ps ...predicate.Priority) *PriorityUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PriorityUpdate) SetName(s string) *PriorityUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PriorityUpdate) SetNillableName(s *string) *PriorityUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PriorityUpdate) SetCreatedAt(t time.Time) *PriorityUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PriorityUpdate) SetNillableCreatedAt(t *time.Time) *PriorityUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PriorityUpdate) SetUpdatedAt(t time.Time) *PriorityUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// AddTodoIDs adds the "todo" edge to the Todo entity by IDs.
func (pu *PriorityUpdate) AddTodoIDs(ids ...int) *PriorityUpdate {
	pu.mutation.AddTodoIDs(ids...)
	return pu
}

// AddTodo adds the "todo" edges to the Todo entity.
func (pu *PriorityUpdate) AddTodo(t ...*Todo) *PriorityUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTodoIDs(ids...)
}

// Mutation returns the PriorityMutation object of the builder.
func (pu *PriorityUpdate) Mutation() *PriorityMutation {
	return pu.mutation
}

// ClearTodo clears all "todo" edges to the Todo entity.
func (pu *PriorityUpdate) ClearTodo() *PriorityUpdate {
	pu.mutation.ClearTodo()
	return pu
}

// RemoveTodoIDs removes the "todo" edge to Todo entities by IDs.
func (pu *PriorityUpdate) RemoveTodoIDs(ids ...int) *PriorityUpdate {
	pu.mutation.RemoveTodoIDs(ids...)
	return pu
}

// RemoveTodo removes "todo" edges to Todo entities.
func (pu *PriorityUpdate) RemoveTodo(t ...*Todo) *PriorityUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PriorityUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PriorityUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PriorityUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PriorityUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PriorityUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := priority.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PriorityUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := priority.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Priority.name": %w`, err)}
		}
	}
	return nil
}

func (pu *PriorityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(priority.Table, priority.Columns, sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(priority.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(priority.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(priority.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.TodoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   priority.TodoTable,
			Columns: []string{priority.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTodoIDs(); len(nodes) > 0 && !pu.mutation.TodoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   priority.TodoTable,
			Columns: []string{priority.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   priority.TodoTable,
			Columns: []string{priority.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{priority.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PriorityUpdateOne is the builder for updating a single Priority entity.
type PriorityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PriorityMutation
}

// SetName sets the "name" field.
func (puo *PriorityUpdateOne) SetName(s string) *PriorityUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PriorityUpdateOne) SetNillableName(s *string) *PriorityUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PriorityUpdateOne) SetCreatedAt(t time.Time) *PriorityUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PriorityUpdateOne) SetNillableCreatedAt(t *time.Time) *PriorityUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PriorityUpdateOne) SetUpdatedAt(t time.Time) *PriorityUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// AddTodoIDs adds the "todo" edge to the Todo entity by IDs.
func (puo *PriorityUpdateOne) AddTodoIDs(ids ...int) *PriorityUpdateOne {
	puo.mutation.AddTodoIDs(ids...)
	return puo
}

// AddTodo adds the "todo" edges to the Todo entity.
func (puo *PriorityUpdateOne) AddTodo(t ...*Todo) *PriorityUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTodoIDs(ids...)
}

// Mutation returns the PriorityMutation object of the builder.
func (puo *PriorityUpdateOne) Mutation() *PriorityMutation {
	return puo.mutation
}

// ClearTodo clears all "todo" edges to the Todo entity.
func (puo *PriorityUpdateOne) ClearTodo() *PriorityUpdateOne {
	puo.mutation.ClearTodo()
	return puo
}

// RemoveTodoIDs removes the "todo" edge to Todo entities by IDs.
func (puo *PriorityUpdateOne) RemoveTodoIDs(ids ...int) *PriorityUpdateOne {
	puo.mutation.RemoveTodoIDs(ids...)
	return puo
}

// RemoveTodo removes "todo" edges to Todo entities.
func (puo *PriorityUpdateOne) RemoveTodo(t ...*Todo) *PriorityUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTodoIDs(ids...)
}

// Where appends a list predicates to the PriorityUpdate builder.
func (puo *PriorityUpdateOne) Where(ps ...predicate.Priority) *PriorityUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PriorityUpdateOne) Select(field string, fields ...string) *PriorityUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Priority entity.
func (puo *PriorityUpdateOne) Save(ctx context.Context) (*Priority, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PriorityUpdateOne) SaveX(ctx context.Context) *Priority {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PriorityUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PriorityUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PriorityUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := priority.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PriorityUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := priority.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Priority.name": %w`, err)}
		}
	}
	return nil
}

func (puo *PriorityUpdateOne) sqlSave(ctx context.Context) (_node *Priority, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(priority.Table, priority.Columns, sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Priority.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, priority.FieldID)
		for _, f := range fields {
			if !priority.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != priority.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(priority.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(priority.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(priority.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.TodoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   priority.TodoTable,
			Columns: []string{priority.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTodoIDs(); len(nodes) > 0 && !puo.mutation.TodoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   priority.TodoTable,
			Columns: []string{priority.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   priority.TodoTable,
			Columns: []string{priority.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Priority{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{priority.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
