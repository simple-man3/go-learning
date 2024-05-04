package repository

type Authorization interface {
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
func NewRepository() *Repository {
	return &Repository{}
}
