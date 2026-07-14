package repository

import (
	"github.com/moeghifar/mariana/infra"
	"github.com/moeghifar/mariana/module/account/repository/user"
)

type Repositories struct {
	User user.UserRepository
}

func Init(deps *infra.Dependencies) Repositories {
	return Repositories{
		User: user.Init(deps),
	}
}
