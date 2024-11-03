package repository

import (
	"context"
	"server/domain/model"
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

type SearchResult struct {
	Labels     []*model.Label
	Statuses   []*model.Status
	Priorities []*model.Priority
}

type TodoRepositoryInterface interface {
	AllTodos(ctx context.Context, input *Input) ([]*model.Todo, error)
	CreateTodo(ctx context.Context, mt *model.Todo, labelIDs []int) (*model.Todo, error)
	Search(ctx context.Context) (SearchResult, error)
}
