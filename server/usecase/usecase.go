package usecase

import (
	"context"
	"server/repogitory"
)

type usecase struct {
	Repogitory repogitory.RepogitoryInterface
}

type UsecaseInterface interface {
	TodosGet(ctx context.Context) ([]Todo, error)
}

func NewUsecase() UsecaseInterface {
	return &usecase{}
}
