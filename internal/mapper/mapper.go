package mapper

import (
	"wish/internal/dto"
	"wish/internal/model"
)

type Users interface {
	MapperToModel(input dto.UserSingUpInput, passwordHash string) model.User
}

type Mappers struct {
	Users Users
}

func NewMappers() *Mappers {
	return &Mappers{
		Users: NewUserMapper(),
	}
}
