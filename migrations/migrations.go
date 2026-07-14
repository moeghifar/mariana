package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

func Up(ctx context.Context, db *bun.DB) error {
	m := migrate.NewMigrator(db, Migrations)
	if err := m.Init(ctx); err != nil {
		return err
	}
	_, err := m.Migrate(ctx)
	return err
}

func Create(ctx context.Context, name string) error {
	m := migrate.NewMigrator(nil, Migrations)
	_, err := m.CreateSQLMigrations(ctx, name)
	return err
}
