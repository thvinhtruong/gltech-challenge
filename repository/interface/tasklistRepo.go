package repository

import (
	entity "github.com/thvinhtruong/legoha/entities"
)

type TaskListRepository interface {
	Assign(tl *entity.TaskList) error
	ListUsersForOneTodo(tl *entity.TaskList) ([]*entity.User, error)
	Completed(tl *entity.TaskList) error
	Revoke(tl *entity.TaskList) error
	Undo(tl *entity.TaskList) error
}
