package service

import (
	todo_go "todo-go"
	"todo-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_go.User) (int, error)
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

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
