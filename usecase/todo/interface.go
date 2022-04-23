package todo

import (
	entity "github.com/thvinhtruong/legoha/entities"
)

type TodoUseCase interface {
	CreateTodo(title, description string) error
	ListTodos() ([]*entity.Todo, error)
	GetTodoByID(id int) (*entity.Todo, error)
	PatchTodo(id int, t *entity.Todo) error
	DeleteTodo(id int) error
}
