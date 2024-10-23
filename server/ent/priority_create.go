// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/priority"
	"server/ent/todo"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PriorityCreate is the builder for creating a Priority entity.
type PriorityCreate struct {
	config
	mutation *PriorityMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PriorityCreate) SetName(s string) *PriorityCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PriorityCreate) SetID(i int) *PriorityCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (pc *PriorityCreate) AddTodoIDs(ids ...int) *PriorityCreate {
	pc.mutation.AddTodoIDs(ids...)
	return pc
}

// AddTodos adds the "todos" edges to the Todo entity.
func (pc *PriorityCreate) AddTodos(t ...*Todo) *PriorityCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTodoIDs(ids...)
}

// Mutation returns the PriorityMutation object of the builder.
func (pc *PriorityCreate) Mutation() *PriorityMutation {
	return pc.mutation
}

// Save creates the Priority in the database.
func (pc *PriorityCreate) Save(ctx context.Context) (*Priority, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PriorityCreate) SaveX(ctx context.Context) *Priority {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PriorityCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PriorityCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PriorityCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Priority.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := priority.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Priority.name": %w`, err)}
		}
	}
	return nil
}

func (pc *PriorityCreate) sqlSave(ctx context.Context) (*Priority, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PriorityCreate) createSpec() (*Priority, *sqlgraph.CreateSpec) {
	var (
		_node = &Priority{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(priority.Table, sqlgraph.NewFieldSpec(priority.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(priority.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := pc.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   priority.TodosTable,
			Columns: priority.TodosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PriorityCreateBulk is the builder for creating many Priority entities in bulk.
type PriorityCreateBulk struct {
	config
	err      error
	builders []*PriorityCreate
}

// Save creates the Priority entities in the database.
func (pcb *PriorityCreateBulk) Save(ctx context.Context) ([]*Priority, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Priority, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PriorityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PriorityCreateBulk) SaveX(ctx context.Context) []*Priority {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PriorityCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PriorityCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
