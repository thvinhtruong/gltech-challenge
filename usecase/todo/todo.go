package todo

import (
	entity "github.com/thvinhtruong/legoha/entities"
	repository "github.com/thvinhtruong/legoha/repository"
)

type TodoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (s *TodoService) CreateTodo(title, description string) error {
	todo := &entity.Todo{Title: title, Description: description, Completed: false}
	err := s.todoRepo.CreateNewTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) ListTodos() ([]*entity.Todo, error) {
	return s.todoRepo.ListTodos()
}

func (s *TodoService) GetTodoByID(id int) (*entity.Todo, error) {
	return s.todoRepo.GetTodoByID(id)
}

func (s *TodoService) PatchTodo(id int) (*entity.Todo, error) {
	todo, err := s.todoRepo.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	err = s.todoRepo.PatchTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
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
