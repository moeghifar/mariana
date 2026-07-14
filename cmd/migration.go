package cmd

import (
	"context"

	"github.com/moeghifar/libgo/pkg/climd"
	"github.com/moeghifar/mariana/config"
	"github.com/moeghifar/mariana/infra"
	"github.com/moeghifar/mariana/migrations"
)

func migrationCommand(conf config.Config) climd.Command {
	return climd.Command{
		Name:  "db",
		Short: "Database migration operations",
		SubCommands: []climd.SubCommand{
			{
				Name:  "up",
				Short: "Apply all pending migrations",
				Run: func(ctx context.Context, args []string, flags map[string]string) error {
					deps, err := infra.Init(ctx, conf)
					if err != nil {
						return err
					}
					return migrations.Up(ctx, deps.DB)
				},
			},
			{
				Name:  "create",
				Short: "Scaffold a new migration file",
				Flags: []climd.Flag{
					{Name: "name", Short: "n", Usage: "migration name", Required: true},
				},
				Run: func(ctx context.Context, args []string, flags map[string]string) error {
					return migrations.Create(ctx, flags["name"])
				},
			},
		},
	}
}
