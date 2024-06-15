package service

import (
	todo_go "todo-go"
	"todo-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_go.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_go.TodoList) (int, error)
	GetAll(userId int) ([]todo_go.TodoList, error)
	GetById(userId int, listId int) (todo_go.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input todo_go.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item todo_go.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo_go.TodoItem, error)
	GetById(userId, itemId int) (todo_go.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_go.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
		TodoItem:      NewTodoItemService(repository.TodoItem, repository.TodoList),
	}
}
