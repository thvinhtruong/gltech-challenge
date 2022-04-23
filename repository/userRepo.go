package repository

import entity "github.com/thvinhtruong/legoha/entities"

type UserRepository interface {
	CreateUser(u *entity.User) error
	ListUsers() ([]*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	PatchUser(u *entity.User) error
	DeleteUser(u *entity.User) error
	LoginUser(u *entity.User) (*entity.User, error)
}
