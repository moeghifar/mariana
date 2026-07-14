package cmd

import (
	"context"

	"github.com/moeghifar/libgo/pkg/climd"
	"github.com/moeghifar/mariana/config"
	"github.com/moeghifar/mariana/infra"
	"github.com/moeghifar/mariana/lib"
	"github.com/moeghifar/mariana/module/account"
	accountpresenter "github.com/moeghifar/mariana/module/account/presenter"
	"github.com/moeghifar/mariana/module/account/usecase"
)

func serveCommand(conf config.Config) climd.Command {
	return climd.Command{
		Name:  "serve",
		Short: "Start the service",
		Long:  "Boots infra and modules, then serves the selected presenters.",
		Flags: []climd.Flag{
			{Name: "http", Usage: "enable the HTTP presenter"},
			{Name: "grpc", Usage: "enable the gRPC presenter"},
			{Name: "consumer", Usage: "enable the message-consumer presenter"},
		},
		Run: func(ctx context.Context, args []string, flags map[string]string) error {
			lib.Init(lib.LogOptions{Level: conf.Log.Level, Format: conf.Log.Format})

			deps, err := infra.Init(ctx, conf)
			if err != nil {
				return err
			}
			defer deps.DB.Close()

			accountModule := account.Init(deps)
			ucs := usecase.Usecases{
				MyUsecase: accountModule.Usecases.MyUsecase,
			}

			_, http := flags["http"]
			_, grpc := flags["grpc"]
			_, consumer := flags["consumer"]

			return accountpresenter.Init(ctx, conf, ucs, accountpresenter.Enable{
				HTTP:     http,
				GRPC:     grpc,
				Consumer: consumer,
			})
		},
	}
}
