package usecase

import (
	"context"
)

type DeleteTodo struct {
	ID int
}

func (u *usecase) DeleteTodo(ctx context.Context, todo DeleteTodo) (int, error) {
	return u.TodoRepositoryInterface.DeleteTodo(ctx, todo.ID)
}
