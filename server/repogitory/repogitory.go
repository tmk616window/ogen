package repogitory

import (
	"context"
	"server/ogen"
)

type repogitory struct{}

type RepogitoryInterface interface {
	GetUsers(ctx context.Context) ([]ogen.Todo, error)
}

func NewRepogitory() RepogitoryInterface {
	return &repogitory{}
}

func (r *repogitory) GetUsers(ctx context.Context) ([]ogen.Todo, error) {
	return []ogen.Todo{
		{
			ID:    "1",
			Title: "title",
		},
	}, nil
}
