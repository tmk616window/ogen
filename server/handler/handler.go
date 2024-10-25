package handler

import (
	"context"
	"server/ogen"
	"server/usecase"
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

	return []ogen.Todo{
		{
			ID: todos[0].ID,
		},
	}, nil
}
