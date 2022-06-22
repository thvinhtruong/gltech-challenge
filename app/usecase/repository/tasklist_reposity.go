package repository

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type TaskListRepository interface {
	Assign(tl *entity.TodoUser) error
	ListUsersForOneTodo(tl *entity.TodoUser) ([]*entity.User, error)
	Completed(tl *entity.TodoUser) error
	Revoke(tl *entity.TodoUser) error
	Undo(tl *entity.TodoUser) error
}
