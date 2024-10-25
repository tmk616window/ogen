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

func (h *handler) TodosGet(ctx context.Context) ([]ogen.Todo, error) {
	todos, err := h.Usecase.TodosGet(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Map(todos, func(todo *usecase.Todo, _ int) ogen.Todo {
		return ogen.Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: ogen.OptString{Value: todo.Description},
		}
	}), nil
}
