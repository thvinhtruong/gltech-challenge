package gormdb

import (
	"context"

	"github.com/jinzhu/gorm"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// add a user
func (r *Repository) CreateUser(ctx context.Context, user entity.User) error {
	setUserRole(user, "user")
	err := r.getTx(ctx).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

// get all user
func (r *Repository) ListUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	err := r.getTx(ctx).Find(users).Error
	if err != nil {
		return nil, err
	}

	result := make([]entity.User, 0, len(users))
	for _, user := range users {
		result = append(result, newUser(user))
	}

	return result, nil
}

// get user by id
func (r *Repository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User

	err := r.getTx(ctx).First(&user, id).Error
	if err != nil {
		return user, err
	}

	return newUser(user), nil
}

// get user by username
func (r *Repository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := r.getTx(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return newUser(user), nil
}

// update user infor
func (r *Repository) PatchUser(ctx context.Context, id int, u entity.User) error {
	var user entity.User
	err := r.getTx(ctx).First(&user, id).Model(&user).Update(u).Error
	if err != nil {
		return err
	}
	return nil
}

// delete user by id
func (r *Repository) DeleteUser(ctx context.Context, u entity.User) error {
	err := r.getTx(ctx).First(&u, u.ID).Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) LoginUser(ctx context.Context, u entity.User) (entity.User, error) {
	var user entity.User
	err := r.getTx(ctx).Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return user, err
	}

	return newUser(user), nil
}
