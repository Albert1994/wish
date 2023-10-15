package service

import (
	"errors"
	"wish/internal/dto"
	"wish/internal/helpers"
	"wish/internal/mapper"
	"wish/internal/repository"
	"wish/pkg/auth"
	"wish/pkg/hash"
)

type UserService struct {
	repo         repository.Users
	tokenManager auth.TokenManager
	Hasher       hash.PasswordHasher
	Mapper       mapper.Users
}

func NewUsersService(repo repository.Users, tokenManager auth.TokenManager, hasher hash.PasswordHasher, mapper mapper.Users) *UserService {
	return &UserService{
		repo:         repo,
		tokenManager: tokenManager,
		Hasher:       hasher,
		Mapper:       mapper,
	}
}

func (s *UserService) SingUp(input dto.UserSingUpInput) error {
	passwordHash, err := s.Hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := s.Mapper.MapperToModel(input, passwordHash)

	if err := s.repo.CreateUser(user); err != nil {
		if errors.Is(err, helpers.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	return nil
}
