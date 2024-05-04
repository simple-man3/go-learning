package service

import "example/hello/pkg/repository"

type Authorization interface {
	Auth(email string, password string) bool
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
	return &Service{}
}
