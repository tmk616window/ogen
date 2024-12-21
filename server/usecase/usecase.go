package usecase

import (
	"context"
	"server/domain/repository"
)

type usecase struct {
	TodoRepositoryInterface repository.TodoRepositoryInterface
}

type UsecaseInterface interface {
	TodosGet(ctx context.Context, input *Input) (*TodosGet, error)
	CreateTodo(ctx context.Context, t *CreateTodo, labelIDs []int) (*Todo, error)
	Search(ctx context.Context) (*SearchResult, error)
	DeleteTodo(ctx context.Context, t DeleteTodo) (int, error)
}

func NewUsecase(rtri repository.TodoRepositoryInterface) UsecaseInterface {
	return &usecase{rtri}
}
