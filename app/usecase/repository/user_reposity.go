package repository

import (
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type UserRepository interface {
	CreateUser(entityUser *entity.User) error
	ListUsers() ([]*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	PatchUser(id int, u *entity.User) error
	DeleteUser(u *entity.User) error
	LoginUser(user *entity.User) (*entity.User, error)
}
