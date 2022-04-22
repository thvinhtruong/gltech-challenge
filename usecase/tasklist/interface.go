package tasklist

import (
	entity "github.com/thvinhtruong/legoha/entities"
)

type TaskListUseCase interface {
	AssignTask(user_id, todo_id int) error
	ListUsersTasks(user_id int) ([]*entity.User, error)
	CompletedTask(user_id, todo_id int) error
	RevokeTask(user_id, todo_id int) error
	UndoTask(user_id, todo_id int) error
}
