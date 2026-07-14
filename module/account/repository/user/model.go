package user

import (
	"github.com/moeghifar/mariana/module/account/entity"
	"github.com/uptrace/bun"
)

type Model struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64  `bun:"id,pk,autoincrement"`
	Name  string `bun:"name"`
	Email string `bun:"email"`
}

func (m Model) ToEntity() entity.User {
	return entity.User{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}

func UserFromEntity(in entity.User) Model {
	return Model{
		ID:    in.ID,
		Name:  in.Name,
		Email: in.Email,
	}
}
