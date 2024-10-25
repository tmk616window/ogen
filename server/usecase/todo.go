package usecase

import (
	"context"
	"server/ent"

	"github.com/samber/lo"
)

type Todo struct {
	ID          int
	Title       string
	Description string
}

func (u *usecase) TodosGet(ctx context.Context) ([]*Todo, error) {
	todos, err := u.db.AllTodos(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Map(todos, func(todo *ent.Todo, _ int) *Todo {
		return &Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
		}
	}), nil
}
