package repository

import (
	"TO_DO_List"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TO_DO_List.User) (int, error)
	GetUser(username, password string) (TO_DO_List.User, error)
}

type TodoList interface {
	Create(userId int, list TO_DO_List.TodoList) (int, error)
	GetAll(id int) ([]TO_DO_List.TodoList, error)
	GetById(userId, listId int) (TO_DO_List.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TO_DO_List.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item TO_DO_List.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TO_DO_List.TodoItem, error)
	GetById(userId, itemId int) (TO_DO_List.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input TO_DO_List.UpdateItemInput) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
