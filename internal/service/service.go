package service

import "wish/internal/repository"

type Users interface {
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{}
}
