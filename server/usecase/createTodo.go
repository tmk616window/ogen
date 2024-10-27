package usecase

import (
	"context"
	"server/db"
	"server/ent"

	"github.com/samber/lo"
)

type CreateTodoInput struct {
	Title       string
	Description string
	LabelIDs    []int
	PriorityID  int
	StatusID    int
}

func (u *usecase) CreateTodo(ctx context.Context, input *CreateTodoInput) (*Todo, error) {
	todo, err := u.db.CreateTodo(ctx, &db.CreateTodoInput{
		Title:       input.Title,
		Description: input.Description,
		LabelsID:    input.LabelIDs,
		StatusID:    input.StatusID,
		PriorityID:  input.PriorityID,
	})
	if err != nil {
		return nil, err
	}

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
	}, nil
}
