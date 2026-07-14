package account

import (
	"github.com/moeghifar/mariana/infra"
	"github.com/moeghifar/mariana/module/account/repository"
	"github.com/moeghifar/mariana/module/account/service"
	"github.com/moeghifar/mariana/module/account/usecase"
)

type Module struct {
	Repositories repository.Repositories
	Services     service.Services
	Usecases     usecase.Usecases
}

func Init(deps *infra.Dependencies) Module {
	repos := repository.Init(deps)
	svcs := service.Init(repos)
	ucs := usecase.Init(svcs)

	return Module{
		Repositories: repos,
		Services:     svcs,
		Usecases:     ucs,
	}
}
