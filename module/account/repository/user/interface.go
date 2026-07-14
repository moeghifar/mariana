package user

import (
	"context"

	"github.com/moeghifar/mariana/module/account/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (entity.User, error)
	Create(ctx context.Context, in entity.User) error
}
