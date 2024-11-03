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

func (h *handler) TodosGet(ctx context.Context, params ogen.TodosGetParams) ([]ogen.Todo, error) {

	todos, err := h.Usecase.TodosGet(ctx, &usecase.Input{
		Limit:  params.Limit.Value,
		Offset: params.Offset.Value,
		WhereInput: usecase.WhereInput{
			Title:       params.WhereTodoInput.Value.Title.Value,
			Description: params.WhereTodoInput.Value.Description.Value,
			// LabelIDs:    params.WhereTodoInput.Value.Labels,
			PriorityID: params.WhereTodoInput.Value.PriorityID.Value,
			StatusID:   params.WhereTodoInput.Value.StatusID.Value,
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

func (h *handler) TodoPost(ctx context.Context, req *ogen.CreateTodoInput) (*ogen.CreateTodoResponse, error) {
	todo, err := h.Usecase.CreateTodo(ctx, &usecase.CreateTodo{
		Title:       req.Title,
		Description: req.Description,
		StatusID:    req.StatusID,
		PriorityID:  req.PriorityID,
	}, req.LabelIDs)
	if err != nil {
		return nil, h.NewError(ctx, err)
	}

	return &ogen.CreateTodoResponse{
		ID: todo.ID,
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
