package usecase

import (
	"context"
	"fmt"
	"server/db"
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
		PriorityID:  input.PriorityID,
		StatusID:    input.StatusID,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(todo)

	return nil, nil
}
