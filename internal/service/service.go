package service

import (
	"wish/internal/dto"
	"wish/internal/mapper"
	"wish/internal/repository"
	"wish/pkg/auth"
	"wish/pkg/hash"
)

type Users interface {
	SingUp(input dto.UserSingUpInput) error
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager auth.TokenManager
	Hasher       hash.PasswordHasher
	Mapper       mapper.Mappers
}

func NewServices(deps Deps) *Services {
	userService := NewUsersService(deps.Repos.Users, deps.TokenManager, deps.Hasher, deps.Mapper.Users)
	return &Services{
		Users: userService,
	}
}
