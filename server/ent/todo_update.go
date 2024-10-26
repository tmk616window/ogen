// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/predicate"
	"server/ent/priority"
	"server/ent/status"
	"server/ent/todo"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TodoUpdate) SetTitle(s string) *TodoUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableTitle(s *string) *TodoUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TodoUpdate) SetDescription(s string) *TodoUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableDescription(s *string) *TodoUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TodoUpdate) ClearDescription() *TodoUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetName sets the "name" field.
func (tu *TodoUpdate) SetName(s string) *TodoUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableName(s *string) *TodoUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetFinishedAt sets the "finished_at" field.
func (tu *TodoUpdate) SetFinishedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetFinishedAt(t)
	return tu
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableFinishedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetFinishedAt(*t)
	}
	return tu
}

// SetPriorityID sets the "priority_id" field.
func (tu *TodoUpdate) SetPriorityID(i int) *TodoUpdate {
	tu.mutation.SetPriorityID(i)
	return tu
}

// SetNillablePriorityID sets the "priority_id" field if the given value is not nil.
func (tu *TodoUpdate) SetNillablePriorityID(i *int) *TodoUpdate {
	if i != nil {
		tu.SetPriorityID(*i)
	}
	return tu
}

// SetStatusID sets the "status_id" field.
func (tu *TodoUpdate) SetStatusID(i int) *TodoUpdate {
	tu.mutation.SetStatusID(i)
	return tu
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableStatusID(i *int) *TodoUpdate {
	if i != nil {
		tu.SetStatusID(*i)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TodoUpdate) SetCreatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableCreatedAt(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TodoUpdate) SetUpdatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetPriority sets the "priority" edge to the Priority entity.
func (tu *TodoUpdate) SetPriority(p *Priority) *TodoUpdate {
	return tu.SetPriorityID(p.ID)
}

// SetStatus sets the "status" edge to the Status entity.
func (tu *TodoUpdate) SetStatus(s *Status) *TodoUpdate {
	return tu.SetStatusID(s.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearPriority clears the "priority" edge to the Priority entity.
func (tu *TodoUpdate) ClearPriority() *TodoUpdate {
	tu.mutation.ClearPriority()
	return tu
}

// ClearStatus clears the "status" edge to the Status entity.
func (tu *TodoUpdate) ClearStatus() *TodoUpdate {
	tu.mutation.ClearStatus()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TodoUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := todo.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := todo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Todo.title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Name(); ok {
		if err := todo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Todo.name": %w`, err)}
		}
	}
	if tu.mutation.PriorityCleared() && len(tu.mutation.PriorityIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Todo.priority"`)
	}
	if tu.mutation.StatusCleared() && len(tu.mutation.StatusIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Todo.status"`)
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(todo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(todo.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(todo.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(todo.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.FinishedAt(); ok {
		_spec.SetField(todo.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.PriorityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.PriorityTable,
			Columns: []string{todo.PriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.PriorityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.PriorityTable,
			Columns: []string{todo.PriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.StatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.StatusTable,
			Columns: []string{todo.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.StatusTable,
			Columns: []string{todo.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetTitle sets the "title" field.
func (tuo *TodoUpdateOne) SetTitle(s string) *TodoUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableTitle(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TodoUpdateOne) SetDescription(s string) *TodoUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableDescription(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TodoUpdateOne) ClearDescription() *TodoUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TodoUpdateOne) SetName(s string) *TodoUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableName(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetFinishedAt sets the "finished_at" field.
func (tuo *TodoUpdateOne) SetFinishedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetFinishedAt(t)
	return tuo
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableFinishedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetFinishedAt(*t)
	}
	return tuo
}

// SetPriorityID sets the "priority_id" field.
func (tuo *TodoUpdateOne) SetPriorityID(i int) *TodoUpdateOne {
	tuo.mutation.SetPriorityID(i)
	return tuo
}

// SetNillablePriorityID sets the "priority_id" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillablePriorityID(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetPriorityID(*i)
	}
	return tuo
}

// SetStatusID sets the "status_id" field.
func (tuo *TodoUpdateOne) SetStatusID(i int) *TodoUpdateOne {
	tuo.mutation.SetStatusID(i)
	return tuo
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableStatusID(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetStatusID(*i)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TodoUpdateOne) SetCreatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableCreatedAt(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TodoUpdateOne) SetUpdatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetPriority sets the "priority" edge to the Priority entity.
func (tuo *TodoUpdateOne) SetPriority(p *Priority) *TodoUpdateOne {
	return tuo.SetPriorityID(p.ID)
}

// SetStatus sets the "status" edge to the Status entity.
func (tuo *TodoUpdateOne) SetStatus(s *Status) *TodoUpdateOne {
	return tuo.SetStatusID(s.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearPriority clears the "priority" edge to the Priority entity.
func (tuo *TodoUpdateOne) ClearPriority() *TodoUpdateOne {
	tuo.mutation.ClearPriority()
	return tuo
}

// ClearStatus clears the "status" edge to the Status entity.
func (tuo *TodoUpdateOne) ClearStatus() *TodoUpdateOne {
	tuo.mutation.ClearStatus()
	return tuo
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TodoUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := todo.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := todo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Todo.title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Name(); ok {
		if err := todo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Todo.name": %w`, err)}
		}
	}
	if tuo.mutation.PriorityCleared() && len(tuo.mutation.PriorityIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Todo.priority"`)
	}
	if tuo.mutation.StatusCleared() && len(tuo.mutation.StatusIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Todo.status"`)
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(todo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(todo.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(todo.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(todo.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.FinishedAt(); ok {
		_spec.SetField(todo.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.PriorityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.PriorityTable,
			Columns: []string{todo.PriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.PriorityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.PriorityTable,
			Columns: []string{todo.PriorityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.StatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.StatusTable,
			Columns: []string{todo.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.StatusTable,
			Columns: []string{todo.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
