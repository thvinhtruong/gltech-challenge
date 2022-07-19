package repository

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type TodoRepository interface {
	CreateNewTodo(t *entity.Todo) error
	ListTodos() ([]*entity.Todo, error)
	GetTodoByID(id int) (*entity.Todo, error)
	PatchTodo(id int, t *entity.Todo) error
	DeleteTodo(todo *entity.Todo) error
}
