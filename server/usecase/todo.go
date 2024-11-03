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

func (u *usecase) TodosGet(ctx context.Context, input *Input) ([]*Todo, error) {
	todos, err := u.TodoRepositoryInterface.AllTodos(ctx, &repository.Input{
		Limit:    input.Limit,
		Offset:   input.Offset,
		LabelIDs: input.LabelIDs,
		WhereInput: repository.WhereInput{
			Title:       input.WhereInput.Title,
			Description: input.WhereInput.Description,
			StatusID:    input.WhereInput.StatusID,
		},
	})
	if err != nil {
		return nil, err
	}

	return lo.Map(todos, func(todo *ent.Todo, _ int) *Todo {
		return &Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Labels: lo.Map(todo.Edges.Labels, func(label *ent.Label, _ int) Label { // Corrected function signature
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
	}), nil
}
