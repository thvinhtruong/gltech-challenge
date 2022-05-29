package tasklistservice

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type TaskService struct {
	taskRepo TaskListRepository
}

func NewTaskService(taskRepo TaskListRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) AssignTask(user_id, todo_id int) error {
	t := entity.TaskList{UserID: user_id, TodoID: todo_id, Finished: false}
	err := s.taskRepo.Assign(&t)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) ListUsersTasks(user_id int) ([]*entity.User, error) {
	t := entity.TaskList{UserID: user_id, Finished: false}
	users, err := s.taskRepo.ListUsersForOneTodo(&t)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *TaskService) CompletedTask(user_id, todo_id int) error {
	t := entity.TaskList{UserID: user_id, TodoID: todo_id, Finished: false}
	err := s.taskRepo.Completed(&t)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) RevokeTask(user_id, todo_id int) error {
	t := entity.TaskList{UserID: user_id, TodoID: todo_id, Finished: false}
	err := s.taskRepo.Revoke(&t)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) UndoTask(user_id, todo_id int) error {
	t := entity.TaskList{UserID: user_id, TodoID: todo_id, Finished: false}
	err := s.taskRepo.Undo(&t)
	if err != nil {
		return err
	}
	return nil
}
