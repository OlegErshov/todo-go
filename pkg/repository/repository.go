package repository

import (
	"github.com/jmoiron/sqlx"
	todo_go "todo-go"
)

type Authorization interface {
	CreateUser(user todo_go.User) (int, error)
	GetUser(username, password string) (todo_go.User, error)
}

type TodoList interface {
	Create(userId int, list todo_go.TodoList) (int, error)
	GetAll(userId int) ([]todo_go.TodoList, error)
	GetById(userId int, listId int) (todo_go.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input todo_go.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo_go.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo_go.TodoItem, error)
	GetById(userId, itemId int) (todo_go.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		TodoList:      NewTodoListRepository(db),
		TodoItem:      NewTodoItemRepository(db),
	}
}
