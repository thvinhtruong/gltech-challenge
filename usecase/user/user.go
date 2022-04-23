package user

import (
	"errors"

	entity "github.com/thvinhtruong/legoha/entities"
	repository "github.com/thvinhtruong/legoha/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(name, username, password string) error {
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

func (s *UserService) ListUsers() ([]*entity.User, error) {
	return s.userRepo.ListUsers()
}

func (s *UserService) GetUserById(id int) (*entity.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) PatchUser(id int) error {
	u, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	err = s.userRepo.PatchUser(u)
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
