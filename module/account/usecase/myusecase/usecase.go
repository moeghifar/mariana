package myusecase

import (
	"context"
	"fmt"

	"github.com/moeghifar/mariana/module/account/entity"
	"github.com/moeghifar/mariana/module/account/service/account"
)

type useCase struct {
	account account.AccountService
}

func Init(svc account.AccountService) *useCase {
	return &useCase{account: svc}
}

func (uc *useCase) GetByID(ctx context.Context, id int64) (entity.User, error) {
	u, err := uc.account.GetByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("account.GetByID: %w", err)
	}
	return u, nil
}
