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
	CreateTodo(ctx context.Context, t *CreateTodo, labelIDs []int) (*Todo, error)
	Search(ctx context.Context) (*SearchResult, error)
}

func NewUsecase(rtri repository.TodoRepositoryInterface) UsecaseInterface {
	return &usecase{rtri}
}
