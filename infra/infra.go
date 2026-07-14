package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/moeghifar/mariana/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Dependencies struct {
	DB *bun.DB
}

func Init(ctx context.Context, conf config.Config) (*Dependencies, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.Database.DSN)))
	sqldb.SetMaxOpenConns(conf.Database.MaxPoolSize)
	sqldb.SetMaxIdleConns(conf.Database.MaxPoolSize)

	db := bun.NewDB(sqldb, pgdialect.New())
	if err := sqldb.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("connect database: %w", err)
	}

	return &Dependencies{DB: db}, nil
}
