package main

import (
	"github.com/moeghifar/mariana/cmd"
	"github.com/moeghifar/mariana/config"
)

// AppVersion is injected at build time:
//
// go build -ldflags "-X main.AppVersion=$(git describe --tags --always)" -o app .
var AppVersion string = "local"

func main() {
	conf := config.Load()
	conf.AppVersion = AppVersion
	cmd.Init(conf)
}
