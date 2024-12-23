package db

import (
	"context"
	"server/domain/model"
	"server/domain/repository"
	"server/ent"
	"server/ent/label"
	"server/ent/predicate"
	"server/ent/priority"
	"server/ent/status"
	"server/ent/todo"
)

func (c *client) AllTodos(ctx context.Context, input *repository.Input) (*repository.TodoGet, error) {
	todoWhere := []predicate.Todo{
		columnFuzzySearch(todo.FieldDescription, input.WhereInput.Description),
		columnFuzzySearch(todo.FieldTitle, input.WhereInput.Title),
	}

	if input.WhereInput.StatusID != 0 {
		todoWhere = append(todoWhere, todo.HasStatusWith(status.ID(input.WhereInput.StatusID)))
	}

	if input.WhereInput.PriorityID != 0 {
		todoWhere = append(todoWhere, todo.HasPriorityWith(priority.ID(input.WhereInput.PriorityID)))
	}

	if len(input.LabelIDs) > 0 {
		for _, labelValue := range input.LabelIDs {
			todoWhere = append(todoWhere, todo.HasLabelsWith(label.ID(labelValue)))
		}
	}
	todos, err := c.client.Todo.
		Query().
		WithLabels().
		WithPriority().
		WithStatus().
		Limit(input.Limit).
		Offset(input.Offset).
		Where(
			todo.And(
				todoWhere...,
			),
		).
		Order(ent.Desc("created_at")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	count, err := c.client.Todo.
		Query().
		Where(todo.And(
			todoWhere...,
		)).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	return &repository.TodoGet{
		Todos:     todos,
		PageCount: count,
	}, nil
}

func (c *client) CreateTodo(ctx context.Context, mt *model.Todo, labelIDs []int) (*model.Todo, error) {
	tx, err := c.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	todo, err := tx.Todo.
		Create().
		SetTitle(mt.Title).
		SetDescription(mt.Description).
		AddLabelIDs(labelIDs...).
		SetStatusID(mt.StatusID).
		SetPriorityID(mt.PriorityID).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return todo, nil
}

func (c *client) Search(ctx context.Context) (repository.SearchResult, error) {
	labels, err := c.client.Label.Query().All(ctx)
	if err != nil {
		return repository.SearchResult{}, err
	}

	status, err := c.client.Status.Query().All(ctx)
	if err != nil {
		return repository.SearchResult{}, err
	}

	priorities, err := c.client.Priority.Query().All(ctx)
	if err != nil {
		return repository.SearchResult{}, err
	}

	return repository.SearchResult{
		Labels:     labels,
		Statuses:   status,
		Priorities: priorities,
	}, nil
}

func (c *client) DeleteTodo(ctx context.Context, id int) (int, error) {
	tx, err := c.client.Tx(ctx)
	if err != nil {
		return 0, err
	}

	err = tx.Todo.UpdateOneID(id).ClearLabels().Exec(ctx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Todo.DeleteOneID(id).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return id, nil
}
