package usecase

import (
	"context"
	"server/domain/repository"
	"server/ent"
	"time"

	"github.com/samber/lo"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Labels      []Label
	CreatedAt   time.Time
	FinishedAt  time.Time
	Priority    Priority
	Status      Status
}

type Priority struct {
	ID   int
	Name string
}

type Status struct {
	ID    int
	Value string
}

type Label struct {
	ID    int
	Value string
}

type Input struct {
	Limit      int
	Offset     int
	LabelIDs   []int
	WhereInput WhereInput
}

type WhereInput struct {
	Title       string
	Description string
	PriorityID  int
	StatusID    int
}

type TodosGet struct {
	Todos     []*Todo
	PageCount int
}

func (u *usecase) TodosGet(ctx context.Context, input *Input) (*TodosGet, error) {
	todos, err := u.TodoRepositoryInterface.AllTodos(ctx, &repository.Input{
		Limit:    input.Limit,
		Offset:   input.Offset,
		LabelIDs: input.LabelIDs,
		WhereInput: repository.WhereInput{
			Title:       input.WhereInput.Title,
			Description: input.WhereInput.Description,
			PriorityID:  input.WhereInput.PriorityID,
			StatusID:    input.WhereInput.StatusID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &TodosGet{
		Todos: lo.Map(todos.Todos, func(todo *ent.Todo, _ int) *Todo {
			return &Todo{
				ID:          todo.ID,
				Title:       todo.Title,
				Description: todo.Description,
				Labels: lo.Map(todo.Edges.Labels, func(label *ent.Label, _ int) Label {
					return Label{
						ID:    label.ID,
						Value: label.Value,
					}
				}),
				CreatedAt:  todo.CreatedAt,
				FinishedAt: todo.FinishedAt,
				Priority: Priority{
					ID:   todo.Edges.Priority.ID,
					Name: todo.Edges.Priority.Name,
				},
				Status: Status{
					ID:    todo.Edges.Status.ID,
					Value: todo.Edges.Status.Value,
				},
			}
		}),
		PageCount: todos.PageCount,
	}, nil
}
