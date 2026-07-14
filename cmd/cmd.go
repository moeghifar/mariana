package cmd

import (
	"github.com/moeghifar/libgo/pkg/climd"
	"github.com/moeghifar/mariana/config"
)

func Init(conf config.Config) {
	app := climd.AppConfig{
		Name:        conf.App.Name,
		Version:     conf.AppVersion,
		Description: "modular monolith service entrypoint",
		Commands: []climd.Command{
			serveCommand(conf),
			migrationCommand(conf),
			cliCommand(conf),
		},
	}

	climd.Execute(app)
}
