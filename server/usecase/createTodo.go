package usecase

import (
	"context"
	"server/domain/model"
	"server/ent"

	"github.com/samber/lo"
)

type CreateTodo struct {
	Title       string
	Description string
	PriorityID  int
	StatusID    int
}

func (u *usecase) CreateTodo(ctx context.Context, t *CreateTodo, labelIDs []int) (*Todo, error) {
	todo, err := u.TodoRepositoryInterface.CreateTodo(ctx, &model.Todo{
		Title:       t.Title,
		Description: t.Description,
		StatusID:    t.StatusID,
		PriorityID:  t.PriorityID,
	},
		labelIDs)
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
