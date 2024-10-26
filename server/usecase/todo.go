package usecase

import (
	"context"
	"server/ent"
	"time"

	"github.com/samber/lo"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	FinishedAt  time.Time
	Priority    priority
	Status      status
}

type priority struct {
	ID   int
	Name string
}

type status struct {
	ID    int
	Value string
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
			CreatedAt:   todo.CreatedAt,
			FinishedAt:  todo.FinishedAt,
			Priority: priority{
				ID:   todo.Edges.Priority.ID,
				Name: todo.Edges.Priority.Name,
			},
			Status: status{
				ID:    todo.Edges.Status.ID,
				Value: todo.Edges.Status.Value,
			},
		}
	}), nil
}
