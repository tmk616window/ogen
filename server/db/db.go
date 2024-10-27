package db

import (
	"context"
	"database/sql"
	"fmt"

	"server/config"
	"server/ent"
	"server/ent/predicate"
	"server/ent/status"
	"server/ent/todo"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type ClientInterface interface {
	AllTodos(ctx context.Context, input *Input) ([]*ent.Todo, error)
}

type client struct {
	client *ent.Client
}

type Input struct {
	Limit      int
	Offset     int
	WhereInput WhereInput
}

type WhereInput struct {
	Title       string
	Description string
	Status      string
}

func New(c config.Database) (ClientInterface, error) {
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s", c.User, c.Password, c.Host, c.Name)

	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	if db.Ping() != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	return &client{
		client: ent.NewClient(ent.Driver(drv)),
	}, nil
}

func (c *client) GetClient() *ent.Client {
	return c.client
}

func (c *client) AllTodos(ctx context.Context, input *Input) ([]*ent.Todo, error) {
	todoWhere := []predicate.Todo{
		columnFuzzySearch(todo.FieldDescription, input.WhereInput.Description),
		columnFuzzySearch(todo.FieldTitle, input.WhereInput.Title),
		todo.HasStatusWith(status.Value("未着手")),
	}

	if input.WhereInput.Status != "" {
		todoWhere = append(todoWhere, todo.HasStatusWith(status.Value("未着手")))
	}

	todos, err := c.client.Todo.
		Query().
		WithPriority().
		WithStatus().
		Limit(input.Limit).
		Offset(input.Offset).
		Where(
			todo.And(
				todoWhere...,
			),
		).
		Order(ent.Desc("created_at")).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func columnFuzzySearch(column string, value string) func(s *entsql.Selector) {
	return func(s *entsql.Selector) {
		s.Where(entsql.Like(column, fmt.Sprintf("%%%s%%", value)))
	}
}
