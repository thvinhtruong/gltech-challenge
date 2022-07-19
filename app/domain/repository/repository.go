package repository

import (
	"context"

	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type Repository interface {
	UserRepository
	TodoRepository
	TodoUserRepository
	TransactionManager
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUserByFlags(ctx context.Context, name string, username string, flag int) ([]entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	PatchUser(ctx context.Context, id int, user entity.User) error
	DeleteUser(ctx context.Context, id int) error
}

type TodoRepository interface {
	CreateNewTodo(ctx context.Context, todo entity.Todo) error
	GetTodosByFlags(ctx context.Context, title string, description string, flag int) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, id int) (entity.Todo, error)
	PatchTodo(ctx context.Context, id int, todo entity.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}

type TodoUserRepository interface {
	Assign(ctx context.Context, td entity.TodoUser) (int, error)
	ListUsersForOneTodo(ctx context.Context, todo_id int) ([]entity.User, error)
	ListTodosForOneUser(ctx context.Context, user_id int) ([]entity.Todo, error)
	Completed(ctx context.Context, todoId int, userID int) error
	Undo(ctx context.Context, todoId int, userId int) error
	Revoke(ctx context.Context, todoId int, userId int) error
}

type TransactionManager interface {
	EnableTx(txFunc func() error) error
}
