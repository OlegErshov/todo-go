package repository

import (
	"github.com/jmoiron/sqlx"
	todo_go "todo-go"
)

type Authorization interface {
	CreateUser(user todo_go.User) (int, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
	}
}
