package interactor

import (
	"context"

	"github.com/jinzhu/copier"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	"github.com/thvinhtruong/legoha/app/domain/repository"
	"github.com/thvinhtruong/legoha/app/usecase/dto"
)

type TodoUseCase struct {
	repo repository.Repository
}

func NewTodoUseCase(repo repository.Repository) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (s *TodoUseCase) PostNewTodo(ctx context.Context, todo dto.Todo) (err error) {
	var record entity.Todo
	err = copier.Copy(&record, &todo)
	if err != nil {
		return err
	}

	err = s.repo.CreateNewTodo(ctx, record)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoUseCase) ShowTodoByID(ctx context.Context, id int) (result dto.Todo, err error) {
	todo, err := s.repo.GetTodoByID(ctx, id)
	if err != nil {
		return result, err
	}

	err = copier.Copy(&result, &todo)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *TodoUseCase) FindTodos(ctx context.Context, todo dto.Todo) (result []dto.Todo, err error) {
	var totalRecords []entity.Todo
	if len(todo.Title) != 0 {
		record, err := s.repo.GetTodosByFlags(ctx, todo.Title, todo.Description, 0)
		if err != nil {
			return result, err
		}

		totalRecords = append(totalRecords, record...)
	}

	if len(todo.Description) != 0 {
		record, err := s.repo.GetTodosByFlags(ctx, todo.Title, todo.Description, 1)
		if err != nil {
			return result, err
		}

		totalRecords = append(totalRecords, record...)
	}

	if err = copier.Copy(&result, &totalRecords); err != nil {
		return result, err
	}

	return result, nil
}

func (s *TodoUseCase) UpdateTodoInfor(ctx context.Context, id int, t dto.Todo) (err error) {
	var record entity.Todo
	err = copier.Copy(&record, &t)
	if err != nil {
		return err
	}

	err = s.repo.PatchTodo(ctx, id, record)
	if err != nil {
		return err
	}
	return nil
}
