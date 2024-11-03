package db

import (
	"context"
	"server/domain/model"
	"server/domain/repository"
	"server/ent"
	"server/ent/label"
	"server/ent/predicate"
	"server/ent/status"
	"server/ent/todo"
)

func (c *client) AllTodos(ctx context.Context, input *repository.Input) ([]*ent.Todo, error) {
	todoWhere := []predicate.Todo{
		columnFuzzySearch(todo.FieldDescription, input.WhereInput.Description),
		columnFuzzySearch(todo.FieldTitle, input.WhereInput.Title),
	}

	if input.WhereInput.StatusID != 0 {
		todoWhere = append(todoWhere, todo.HasStatusWith(status.ID(input.WhereInput.StatusID)))
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

	return todos, nil
}

func (c *client) CreateTodo(ctx context.Context, mt *model.Todo, labelIDs []int) (*ent.Todo, error) {
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
