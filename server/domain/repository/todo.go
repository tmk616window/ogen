package repository

import (
	"context"
	"server/domain/model"
	"server/ent"
)

type Input struct {
	Limit      int
	Offset     int
	LabelIDs   []int
	WhereInput WhereInput
}

type WhereInput struct {
	Title       string
	Description string
	PriorityID  int
	StatusID    int
}

type TodoRepositoryInterface interface {
	AllTodos(ctx context.Context, input *Input) ([]*ent.Todo, error)
	CreateTodo(ctx context.Context, mt *model.Todo, labelIDs []int) (*ent.Todo, error)
}
