package myusecase

import (
	"context"

	"github.com/moeghifar/mariana/module/account/entity"
)

type MyUsecase interface {
	GetByID(ctx context.Context, id int64) (entity.User, error)
}
