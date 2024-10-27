package handler

import (
	"context"
	"server/ogen"
	"server/usecase"

	"github.com/samber/lo"
)

type handler struct {
	Usecase usecase.UsecaseInterface
}

func NewHandler(u usecase.UsecaseInterface) *handler {
	return &handler{
		Usecase: u,
	}
}

func (h *handler) TodosGet(ctx context.Context, req *ogen.TodoInput) ([]ogen.Todo, error) {
	todos, err := h.Usecase.TodosGet(ctx, &usecase.Input{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(todos, func(todo *usecase.Todo, _ int) ogen.Todo {
		return ogen.Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: ogen.OptString{Value: todo.Description},
			CreatedAt:   todo.CreatedAt,
			FinishedAt:  ogen.OptDateTime{Value: todo.FinishedAt},
			Priority: ogen.Priority{
				ID:   todo.Priority.ID,
				Name: todo.Priority.Name,
			},
			Status: ogen.Status{
				ID:    todo.Status.ID,
				Value: todo.Status.Value,
			},
		}
	}), nil
}
