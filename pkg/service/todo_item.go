package service

import (
	"TO_DO_List"
	"TO_DO_List/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}
func (s *TodoItemService) Create(userId, listId int, item TO_DO_List.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err == nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]TO_DO_List.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
