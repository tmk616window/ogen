package usecase

import (
	"context"
)

type Todo struct {
	ID          string
	Title       string
	Description string
}

func (u *usecase) TodosGet(ctx context.Context) ([]Todo, error) {
	todos, err := u.Repogitory.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return []Todo{
		{
			ID:          todos[0].ID,
			Title:       todos[0].Title,
			Description: "description",
		},
	}, nil
}
