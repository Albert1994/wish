package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"wish/internal/model"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user model.User) error {
	query := fmt.Sprintf("INSERT INTO #{usersTable} (name, last_name, email, password, created_by) values ($1, $2, $3, $4, $5) RETURNING ID")

	row := r.db.QueryRow(query, user.Name, user.LastName, user.Email, user.Password, user.CreatedBy)

	if err := row.Scan(); err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetUser(email string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("Select * From #{usersTable} where email=$1")

	err := r.db.Get(&user, query, email)

	return user, err
}
