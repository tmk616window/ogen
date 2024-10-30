package db

import (
	"context"
	"fmt"
	"server/domain/model"
	"server/domain/repository"
	"server/ent"
	"server/ent/label"
	"server/ent/predicate"
	"server/ent/status"
	"server/ent/todo"

	entsql "entgo.io/ent/dialect/sql"
)

func (c *client) AllTodos(ctx context.Context, input *repository.Input) ([]*ent.Todo, error) {
	todoWhere := []predicate.Todo{
		columnFuzzySearch(todo.FieldDescription, input.WhereInput.Description),
		columnFuzzySearch(todo.FieldTitle, input.WhereInput.Title),
	}

	if input.WhereInput.Status != "" {
		todoWhere = append(todoWhere, todo.HasStatusWith(status.Value(input.WhereInput.Status)))
	}

	if len(input.WhereInput.Labels) > 0 {
		for _, labelValue := range input.WhereInput.Labels {
			todoWhere = append(todoWhere, todo.HasLabelsWith(label.Value(labelValue)))
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

func (c *client) CreateTodo(ctx context.Context, td *model.Todo, labelIDs []int) (*ent.Todo, error) {
	tx, err := c.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	t, err := tx.Todo.
		Create().
		SetTitle(td.Title).
		SetDescription(td.Description).
		AddLabelIDs(labelIDs...).
		SetStatusID(td.StatusID).
		SetPriorityID(td.PriorityID).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tt, err := tx.Todo.
		Query().
		Where(
			entsql.FieldEQ(todo.FieldID, t.ID),
		).
		WithStatus().
		WithPriority().
		WithLabels().
		Only(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return tt, nil
}

func columnFuzzySearch(column string, value string) func(s *entsql.Selector) {
	return func(s *entsql.Selector) {
		s.Where(entsql.Like(column, fmt.Sprintf("%%%s%%", value)))
	}
}
