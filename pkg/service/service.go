package service

import (
	"TO_DO_List"
	"TO_DO_List/pkg/repository"
)

type Authorization interface {
	CreateUser(input TO_DO_List.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list TO_DO_List.TodoList) (int, error)
	GetAll(userId int) ([]TO_DO_List.TodoList, error)
	GetById(userId, listId int) (TO_DO_List.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TO_DO_List.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item TO_DO_List.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TO_DO_List.TodoItem, error)
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
