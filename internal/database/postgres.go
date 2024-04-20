package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, event *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}
func (d DBLogger) AfterQuery(ctx context.Context, event *pg.QueryEvent) error {
	formattedQuery, err := event.FormattedQuery()
	if err != nil {
		return err
	}

	fmt.Println(string(formattedQuery))
	return nil
}

func New(opts *pg.Options) *pg.DB {
	orm.SetTableNameInflector(func(s string) string {
		return s
	})
	return pg.Connect(opts)
}
