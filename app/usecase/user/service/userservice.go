package userusecase

import (
	"errors"

	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

type UserService struct {
	userRepo Repository
}

func NewUserService(userRepo Repository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(name, username, password string) error {
	_, err := s.userRepo.GetUserByUsername(username)
	if err == nil {
		return errors.New("user already exists")
	}
	u := entity.User{Name: name, Username: username, Password: password, Role: "user"}
	err = s.userRepo.CreateUser(&u)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) FindAllUsers() ([]*entity.User, error) {
	return s.userRepo.ListUsers()
}

func (s *UserService) FindUserById(id int) (*entity.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) PatchUserInfor(id int, u *entity.User) error {
	err := s.userRepo.PatchUser(id, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	u, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	err = s.userRepo.DeleteUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) LoginUser(username, password string) (*entity.User, error) {
	u, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if u.Password != password {
		return nil, errors.New("wrong password")
	}
	user := entity.User{Username: username, Password: password}
	return s.userRepo.LoginUser(&user)
}
