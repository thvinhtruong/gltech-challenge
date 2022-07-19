package interactor

import (
	"context"

	"github.com/jinzhu/copier"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	"github.com/thvinhtruong/legoha/app/domain/repository"
	"github.com/thvinhtruong/legoha/app/usecase/dto"
)

type UserUseCase struct {
	repo repository.Repository
}

func NewUserUseCase(repo repository.Repository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (s *UserUseCase) RegisterUser(ctx context.Context, user dto.User) (err error) {
	var record entity.User
	err = copier.Copy(&record, &user)
	if err != nil {
		return err
	}

	err = s.repo.CreateUser(ctx, record)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserUseCase) FindUserByID(ctx context.Context, id int) (result dto.User, err error) {
	u, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return result, err
	}
	err = copier.Copy(&result, &u)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *UserUseCase) FindUsers(ctx context.Context, user dto.User) (result []dto.User, err error) {
	var totalRecords []entity.User

	if len(user.Name) != 0 {
		record, err := s.repo.GetUserByFlags(ctx, user.Name, user.Username, 0)
		if err != nil {
			return result, err
		}

		totalRecords = append(totalRecords, record...)
	}
	if len(user.Username) != 0 {
		record, err := s.repo.GetUserByFlags(ctx, user.Name, user.Username, 1)
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

func (s *UserUseCase) PatchUserInfor(ctx context.Context, id int, user dto.User) (err error) {
	var record entity.User
	err = copier.Copy(&record, &user)
	if err != nil {
		return err
	}

	err = s.repo.PatchUser(ctx, id, record)
	if err != nil {
		return err
	}
	return nil
}
