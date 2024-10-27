package usecase

import (
	"context"
	"server/db"
)

type usecase struct {
	db db.ClientInterface
}

type UsecaseInterface interface {
	TodosGet(ctx context.Context, input *Input) ([]*Todo, error)
	CreateTodo(ctx context.Context, input *CreateTodoInput) (*Todo, error)
}

func NewUsecase(dbc db.ClientInterface) UsecaseInterface {
	return &usecase{dbc}
}
