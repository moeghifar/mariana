package presenter

import (
	"context"
	"log/slog"

	"github.com/moeghifar/mariana/config"
	"github.com/moeghifar/mariana/module/account/usecase"
)

type Enable struct {
	HTTP     bool
	GRPC     bool
	Consumer bool
}

func Init(ctx context.Context, conf config.Config, ucs usecase.Usecases, enable Enable) error {
	if enable.HTTP {
		slog.InfoContext(ctx, "http presenter enabled", "app", conf.App.Name)
	}
	if enable.GRPC {
		slog.InfoContext(ctx, "grpc presenter enabled", "app", conf.App.Name)
	}
	if enable.Consumer {
		slog.InfoContext(ctx, "consumer presenter enabled", "app", conf.App.Name)
	}

	return nil
}
