package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedBy time.Time `json:"created_by"`
}
