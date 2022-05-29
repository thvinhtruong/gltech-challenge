package tasklistservice

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type TaskListRepository interface {
	Assign(tl *entity.TaskList) error
	ListUsersForOneTodo(tl *entity.TaskList) ([]*entity.User, error)
	Completed(tl *entity.TaskList) error
	Revoke(tl *entity.TaskList) error
	Undo(tl *entity.TaskList) error
}

type TaskListUseCase interface {
	AssignTask(user_id, todo_id int) error
	ListUsersTasks(user_id int) ([]*entity.User, error)
	CompletedTask(user_id, todo_id int) error
	RevokeTask(user_id, todo_id int) error
	UndoTask(user_id, todo_id int) error
}
