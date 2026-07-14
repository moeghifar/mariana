package cmd

import (
	"context"

	"github.com/moeghifar/libgo/pkg/climd"
	"github.com/moeghifar/mariana/config"
)

func cliCommand(conf config.Config) climd.Command {
	return climd.Command{
		Name:  "cli",
		Short: "Run a one-off CLI command",
		Run: func(ctx context.Context, args []string, flags map[string]string) error {
			return nil
		},
	}
}
