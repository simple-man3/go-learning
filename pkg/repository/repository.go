package repository

import (
	server "example/hello"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user server.User) (int64, error)
	GetUserById()
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository что-то по типу конструктора
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
