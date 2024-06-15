package service

import (
	todo_go "todo-go"
	"todo-go/pkg/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(itemRepo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{itemRepo: itemRepo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId int, listId int, item todo_go.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.itemRepo.Create(listId, item)
}
