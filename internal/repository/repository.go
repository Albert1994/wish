package repository

import "github.com/jmoiron/sqlx"

type Users interface {
}

type Repositories struct {
	Users Users
}

func NewRepository(db *sqlx.DB) *Repositories {
	return &Repositories{}
}
