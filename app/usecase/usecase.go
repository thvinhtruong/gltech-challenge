package usecase

import (
	"context"

	"github.com/thvinhtruong/legoha/app/usecase/dto"
)

type UserService interface {
	RegisterUser(ctx context.Context, user dto.User) (err error)
	FindUserByID(ctx context.Context, id int) (result dto.User, err error)
	FindUsers(ctx context.Context, user dto.User) (result []dto.User, err error)
	PatchUserInfor(ctx context.Context, id int, user dto.User) (err error)
}

type TodoService interface {
	PostNewTodo(ctx context.Context, todo dto.Todo) (err error)
	ShowTodoByID(ctx context.Context, id int) (result dto.Todo, err error)
	FindTodos(ctx context.Context, todo dto.Todo) (result []dto.Todo, err error)
	UpdateTodoInfor(ctx context.Context, id int, todo dto.Todo) (err error)
}

type AdminService interface {
	AssignTask(ctx context.Context, todouser dto.TodoUser) (err error)
	FindAllUsers(ctx context.Context) (result []dto.User, err error)
	FindAllTodos(ctx context.Context) (result []dto.Todo, err error)
	ListUsersForOneTodo(ctx context.Context, todoId int) (results []dto.User, err error)
	CompletedTask(ctx context.Context, user_id int, todo_id int) (err error)
	RevokeTask(ctx context.Context, user_id int, todo_id int) (err error)
	UndoTask(ctx context.Context, user_id int, todo_id int) (err error)
	DeleteUser(ctx context.Context, id int) (err error)
	DeleteTodo(ctx context.Context, id int) (err error)
}
