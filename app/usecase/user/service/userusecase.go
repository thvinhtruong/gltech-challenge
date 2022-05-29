package userusecase

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type Repository interface {
	CreateUser(entityUser *entity.User) error
	ListUsers() ([]*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	PatchUser(id int, u *entity.User) error
	DeleteUser(u *entity.User) error
	LoginUser(user *entity.User) (*entity.User, error)
}

type UserUseCase interface {
	RegisterUser(name, username, password string) error
	FindAllUsers() ([]*entity.User, error)
	FindUserById(id int) (*entity.User, error)
	PatchUserInfor(id int, user *entity.User) error
	DeleteUser(id int) error
	LoginUser(username, password string) (*entity.User, error)
}
