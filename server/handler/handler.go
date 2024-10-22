package handler

import (
	"context"
	"server/ogen"
	"server/usecase"
)

type Handler struct {
	Usecase usecase.UsecaseInterface
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) TodosGet(ctx context.Context) ([]ogen.Todo, error) {
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

func (h *Handler) TodosPost(ctx context.Context, req *ogen.TodoInput) (*ogen.Todo, error) {
	return nil, nil
}
