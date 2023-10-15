package mapper

import (
	"time"
	"wish/internal/dto"
	"wish/internal/model"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (s *UserMapper) MapperToModel(input dto.UserSingUpInput, passwordHash string) model.User {
	return model.User{
		Name:      input.Name,
		LastName:  input.LastName,
		Password:  passwordHash,
		Email:     input.Email,
		CreatedBy: time.Now(),
	}
}
