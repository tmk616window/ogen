package usecase

import (
	"context"
	"server/domain/repository"
)

type usecase struct {
	TodoRepositoryInterface repository.TodoRepositoryInterface
}

type UsecaseInterface interface {
	TodosGet(ctx context.Context, input *Input) ([]*Todo, error)
	CreateTodo(ctx context.Context, todo *CreateTodo, labelIDs []int) (*Todo, error)
}

func NewUsecase(rtri repository.TodoRepositoryInterface) UsecaseInterface {
	return &usecase{rtri}
}
