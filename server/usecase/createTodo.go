package usecase

import (
	"context"
	"server/domain/model"
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
		FinishedAt:  todo.FinishedAt,
	}, nil
}
