package repository

import (
	"github.com/jmoiron/sqlx"
	"wish/internal/model"
)

type Users interface {
	CreateUser(user model.User) error
	GetUser(email string) (model.User, error)
}

type Repositories struct {
	Users Users
}

func NewRepository(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUserPostgres(db),
	}
}
