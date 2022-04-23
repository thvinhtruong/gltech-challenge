package user

import (
	entity "github.com/thvinhtruong/legoha/entities"
)

type UserUseCase interface {
	CreateUser(name, username, password string) error
	ListUsers() ([]*entity.User, error)
	GetUserById(id int) (*entity.User, error)
	PatchUser(id int, user *entity.User) error
	DeleteUser(id int) error
	LoginUser(username, password string) (*entity.User, error)
}
