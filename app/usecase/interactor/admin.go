package interactor

import (
	"context"

	"github.com/jinzhu/copier"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	"github.com/thvinhtruong/legoha/app/domain/repository"
	"github.com/thvinhtruong/legoha/app/usecase/dto"
)

type AdminUseCase struct {
	repo repository.Repository
}

func NewAdminUseCase(repo repository.Repository) *AdminUseCase {
	return &AdminUseCase{repo: repo}
}

func (s *AdminUseCase) AssignTask(ctx context.Context, todouser dto.TodoUser) (err error) {
	record := entity.TodoUser{UserID: todouser.UserID, TodoID: todouser.TodoID, Finished: false}
	_, err = s.repo.Assign(ctx, record)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminUseCase) FindAllUsers(ctx context.Context) (results []dto.User, err error) {
	users, err := s.repo.GetUserByFlags(ctx, "", "", 2)
	if err != nil {
		return results, err
	}

	for _, user := range users {
		var result dto.User
		err = copier.Copy(&result, &user)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (s *AdminUseCase) FindAllTodos(ctx context.Context) (result []dto.Todo, err error) {
	todos, err := s.repo.GetTodosByFlags(ctx, "", "", 2)
	if err != nil {
		return result, err
	}

	for _, todo := range todos {
		var resultdto dto.Todo
		err = copier.Copy(&resultdto, &todo)
		if err != nil {
			return result, err
		}
		result = append(result, resultdto)
	}
	return result, nil
}

func (s *AdminUseCase) ListUsersForOneTodo(ctx context.Context, todoId int) (results []dto.User, err error) {
	var users []entity.User
	users, err = s.repo.ListUsersForOneTodo(ctx, todoId)
	if err != nil {
		return results, err
	}
	for _, user := range users {
		var result dto.User
		err = copier.Copy(&result, &user)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (s *AdminUseCase) CompletedTask(ctx context.Context, user_id int, todo_id int) (err error) {
	err = s.repo.Completed(ctx, todo_id, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminUseCase) RevokeTask(ctx context.Context, user_id int, todo_id int) (err error) {
	err = s.repo.Revoke(ctx, todo_id, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminUseCase) UndoTask(ctx context.Context, user_id int, todo_id int) (err error) {
	err = s.repo.Undo(ctx, todo_id, user_id)
	if err != nil {
		return err
	}
	return nil
}

// Transaction involves
// Delete user, revoke all user todo
func (s *AdminUseCase) DeleteUser(ctx context.Context, id int) (err error) {
	todos, err := s.repo.ListTodosForOneUser(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.EnableTx(func() error {
		err := s.repo.DeleteUser(ctx, id)
		if err != nil {
			return err
		}

		for _, todo := range todos {
			err := s.repo.Revoke(ctx, todo.ID, id)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// transaction involves
// delete a todo, revoke all the assigned users
func (s *AdminUseCase) DeleteTodo(ctx context.Context, id int) (err error) {
	users, err := s.repo.ListUsersForOneTodo(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.EnableTx(func() error {
		err := s.repo.DeleteTodo(ctx, id)
		if err != nil {
			return err
		}

		for _, user := range users {
			err := s.repo.Revoke(ctx, id, user.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
