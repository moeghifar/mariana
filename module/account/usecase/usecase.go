package usecase

import (
	"github.com/moeghifar/mariana/module/account/service"
	"github.com/moeghifar/mariana/module/account/usecase/myusecase"
)

type Usecases struct {
	MyUsecase myusecase.MyUsecase
}

func Init(svcs service.Services) Usecases {
	return Usecases{
		MyUsecase: myusecase.Init(svcs.Account),
	}
}
