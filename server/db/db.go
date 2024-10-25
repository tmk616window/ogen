package db

import (
	"context"
	"database/sql"
	"fmt"

	"server/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type ClientInterface interface {
	AllTodos(ctx context.Context) ([]*ent.Todo, error)
}

type client struct {
	client *ent.Client
}

func New() (ClientInterface, error) {
	databaseUrl := fmt.Sprint("postgresql://user:password@db/db")

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

func (c *client) AllTodos(ctx context.Context) ([]*ent.Todo, error) {
	todos, err := c.client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
