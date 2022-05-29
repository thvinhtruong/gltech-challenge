package todoservice

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

type TodoUseCase interface {
	PostNewTodo(title, description string) error
	ShowAllTodos() ([]*entity.Todo, error)
	ShowTodoByID(id int) (*entity.Todo, error)
	UpdateTodoInfor(id int, t *entity.Todo) error
	DeleteTodo(id int) error
}
