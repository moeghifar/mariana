package account

import (
	"context"
	"fmt"

	"github.com/moeghifar/mariana/module/account/entity"
	"github.com/moeghifar/mariana/module/account/repository/user"
)

type service struct {
	users user.UserRepository
}

func Init(users user.UserRepository) *service {
	return &service{users: users}
}

func (s *service) GetByID(ctx context.Context, id int64) (entity.User, error) {
	u, err := s.users.GetByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("users.GetByID: %w", err)
	}
	return u, nil
}

func (s *service) Create(ctx context.Context, in entity.User) error {
	if err := s.users.Create(ctx, in); err != nil {
		return fmt.Errorf("users.Create: %w", err)
	}
	return nil
}
