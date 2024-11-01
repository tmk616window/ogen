package db

import (
	"database/sql"
	"fmt"

	"server/config"
	"server/domain/repository"
	"server/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type client struct {
	client *ent.Client
}

func New(c config.Database) (repository.TodoRepositoryInterface, error) {
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

func columnFuzzySearch(column string, value string) func(s *entsql.Selector) {
	return func(s *entsql.Selector) {
		s.Where(entsql.Like(column, fmt.Sprintf("%%%s%%", value)))
	}
}
