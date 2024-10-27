package usecase

import (
	"context"
	"server/db"
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

type Input struct {
	Limit  int
	Offset int
}

func (u *usecase) TodosGet(ctx context.Context, input *Input) ([]*Todo, error) {
	todos, err := u.db.AllTodos(ctx, &db.Input{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
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
