package handler

import (
	"context"
	"net/http"
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
		WhereInput: usecase.WhereInput{
			Title:       req.WhereInput.Value.Title.Value,
			Description: req.WhereInput.Value.Description.Value,
			Labels:      req.WhereInput.Value.Labels,
			Status:      req.WhereInput.Value.Status.Value,
		},
	})
	if err != nil {
		return nil, h.NewError(ctx, err)
	}

	return lo.Map(todos, func(todo *usecase.Todo, _ int) ogen.Todo {
		return ogen.Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: ogen.OptString{Value: todo.Description},
			Labels: lo.Map(todo.Labels, func(label usecase.Label, _ int) ogen.Label {
				return ogen.Label{
					ID:    label.ID,
					Value: label.Value,
				}
			}),
			CreatedAt:  todo.CreatedAt,
			FinishedAt: ogen.OptDateTime{Value: todo.FinishedAt},
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

func (h *handler) TodoPost(ctx context.Context, req *ogen.CreateTodoInput) (*ogen.Todo, error) {
	todo, err := h.Usecase.CreateTodo(ctx, &usecase.CreateTodoInput{
		Title:       req.Title.Value,
		Description: req.Description.Value,
		LabelIDs:    req.LabelIDs,
		StatusID:    req.StatusID.Value,
		PriorityID:  req.PriorityID.Value,
	})
	if err != nil {
		return nil, h.NewError(ctx, err)
	}

	return &ogen.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: ogen.OptString{Value: todo.Description},
		Labels: lo.Map(todo.Labels, func(label usecase.Label, _ int) ogen.Label {
			return ogen.Label{
				ID:    label.ID,
				Value: label.Value,
			}
		}),
		Status: ogen.Status{
			ID:    todo.Status.ID,
			Value: todo.Status.Value,
		},
		Priority: ogen.Priority{
			ID:   todo.Priority.ID,
			Name: todo.Priority.Name,
		},
	}, nil
}

func (h *handler) NewError(ctx context.Context, err error) *ogen.ErrorResponseStatusCode {
	return &ogen.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: ogen.ErrorResponse{
			Error: err.Error(),
		},
	}
}
