package service

import (
	"github.com/moeghifar/mariana/module/account/repository"
	"github.com/moeghifar/mariana/module/account/service/account"
)

type Services struct {
	Account account.AccountService
}

func Init(repos repository.Repositories) Services {
	return Services{
		Account: account.Init(repos.User),
	}
}
