package usecase

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type UserUseCase interface {
	RegisterUser(name, username, password string) error
	FindAllUsers() ([]*entity.User, error)
	FindUserById(id int) (*entity.User, error)
	PatchUserInfor(id int, user *entity.User) error
	DeleteUser(id int) error
	LoginUser(username, password string) (*entity.User, error)
}

type TodoUseCase interface {
	PostNewTodo(title, description string) error
	ShowAllTodos() ([]*entity.Todo, error)
	ShowTodoByID(id int) (*entity.Todo, error)
	UpdateTodoInfor(id int, t *entity.Todo) error
	DeleteTodo(id int) error
}

type TaskListUseCase interface {
	AssignTask(user_id, todo_id int) error
	ListUsersTasks(user_id int) ([]*entity.User, error)
	CompletedTask(user_id, todo_id int) error
	RevokeTask(user_id, todo_id int) error
	UndoTask(user_id, todo_id int) error
}
