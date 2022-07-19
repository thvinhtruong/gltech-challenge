package interactor

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	repository "github.com/thvinhtruong/legoha/app/usecase/repository"
)

type TodoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (s *TodoService) PostNewTodo(title, description string) error {
	todo := &entity.Todo{Title: title, Description: description, Completed: false}
	err := s.todoRepo.CreateNewTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) ShowAllTodos() ([]*entity.Todo, error) {
	return s.todoRepo.ListTodos()
}

func (s *TodoService) ShowTodoByID(id int) (*entity.Todo, error) {
	return s.todoRepo.GetTodoByID(id)
}

func (s *TodoService) UpdateTodoInfor(id int, t *entity.Todo) error {
	err := s.todoRepo.PatchTodo(id, t)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) DeleteTodo(id int) error {
	todo, err := s.todoRepo.GetTodoByID(id)
	if err != nil {
		return err
	}
	err = s.todoRepo.DeleteTodo(todo)
	if err != nil {
		return err
	}
	return nil
}
