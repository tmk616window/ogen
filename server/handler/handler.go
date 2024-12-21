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

func (h *handler) TodosGet(ctx context.Context, params ogen.TodosGetParams) (*ogen.TodosGetOK, error) {
	todos, err := h.Usecase.TodosGet(ctx, &usecase.Input{
		Limit:    params.Limit.Value,
		Offset:   params.Offset.Value,
		LabelIDs: params.LabelIDs,
		WhereInput: usecase.WhereInput{
			Title:       params.Title.Value,
			Description: params.Description.Value,
			PriorityID:  params.PriorityID.Value,
			StatusID:    params.StatusID.Value,
		},
	})

	if err != nil {
		return nil, h.NewError(ctx, err)
	}

	return &ogen.TodosGetOK{
		TodoList: lo.Map(todos.Todos, func(todo *usecase.Todo, _ int) ogen.Todo {
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
		}),
		PageCount: todos.PageCount,
	}, nil
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

func (h *handler) SearchGet(ctx context.Context) (*ogen.ResponseSearchTodo, error) {
	search, err := h.Usecase.Search(ctx)
	if err != nil {
		return nil, h.NewError(ctx, err)
	}

	return &ogen.ResponseSearchTodo{
		Labels: lo.Map(search.Labels, func(label *usecase.Label, _ int) ogen.Label {
			return ogen.Label{
				ID:    label.ID,
				Value: label.Value,
			}
		}),
		Status: lo.Map(search.Statuses, func(status *usecase.Status, _ int) ogen.Status {
			return ogen.Status{
				ID:    status.ID,
				Value: status.Value,
			}
		}),
		Priorities: lo.Map(search.Priorities, func(priority *usecase.Priority, _ int) ogen.Priority {
			return ogen.Priority{
				ID:   priority.ID,
				Name: priority.Name,
			}
		}),
	}, nil
}

func (h *handler) TodoIDPut(ctx context.Context, req *ogen.WhereTodoInput, params ogen.TodoIDPutParams) (int, error) {
	return 0, nil
}

func (h *handler) TodoIDDelete(ctx context.Context, params ogen.TodoIDDeleteParams) (int, error) {
	return h.Usecase.DeleteTodo(ctx, usecase.DeleteTodo{ID: params.ID})
}

func (h *handler) NewError(ctx context.Context, err error) *ogen.ErrorResponseStatusCode {
	return &ogen.ErrorResponseStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: ogen.ErrorResponse{
			Error: err.Error(),
		},
	}
}
