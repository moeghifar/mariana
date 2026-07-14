package user

import (
	"context"
	"fmt"

	"github.com/moeghifar/mariana/infra"
	"github.com/moeghifar/mariana/module/account/entity"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func Init(deps *infra.Dependencies) *Repository {
	return &Repository{db: deps.DB}
}

func (r *Repository) GetByID(ctx context.Context, id int64) (entity.User, error) {
	var model Model
	if err := r.db.NewSelect().
		Model(&model).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return entity.User{}, fmt.Errorf("select user: %w", err)
	}
	return model.ToEntity(), nil
}

func (r *Repository) Create(ctx context.Context, in entity.User) error {
	model := UserFromEntity(in)
	if _, err := r.db.NewInsert().Model(&model).Exec(ctx); err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}
