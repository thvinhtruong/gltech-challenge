package gormdb

import (
	"github.com/jinzhu/gorm"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func NewTaskListRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Assign(tl *entity.TodoUser) error {
	t := entity.TodoUser{UserID: tl.UserID, TodoID: tl.TodoID, Finished: false}
	err := r.DB.Create(&t).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListUsersForOneTodo(tl entity.TodoUser) ([]entity.User, error) {
	var users []entity.User

	err := r.DB.Find(&users).Where("id = ?", tl.UserID).Error
	if err != nil {
		return users, err
	}

	result := make([]entity.User, 0, len(users))
	for _, user := range users {
		result = append(result, newUser(user))
	}
	return result, err

}

func (r *Repository) Completed(tl entity.TodoUser) error {
	err := r.DB.Find(&tl).Where("user_id = ? AND todo_id = ?", tl.UserID, tl.TodoID).Error
	if err != nil {
		return err
	}

	setTodoStatus(tl, true)
	err = r.DB.Save(&tl).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Undo(tl entity.TodoUser) error {
	err := r.DB.Find(&tl).Where("user_id = ? AND todo_id = ?", tl.UserID, tl.TodoID).Error
	if err != nil {
		return err
	}

	setTodoStatus(tl, false)
	err = r.DB.Save(&tl).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Revoke(tl *entity.TodoUser) error {
	err := r.DB.Delete(&tl).Error
	if err != nil {
		return err
	}

	return nil
}
