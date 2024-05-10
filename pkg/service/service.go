package service

import (
	server "example/hello"
	"example/hello/pkg/repository"
)

type Authorization interface {
	CreateUser(user server.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func (s *Service) Auth(email string, password string) bool {
	return true
}

// NewService что-то по типу конструктора
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
