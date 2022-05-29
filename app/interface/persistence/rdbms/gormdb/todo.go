package repository

import (
	"github.com/jinzhu/gorm"
	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func NewTodo(todo *entity.Todo) *entity.Todo {
	t := entity.Todo{}
	t.ID = todo.ID
	t.Title = todo.Title
	t.Description = todo.Description
	t.Completed = todo.Completed
	t.CreatedAt = todo.CreatedAt
	return &t
}

func NewTodoRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateNewTodo(t *entity.Todo) error {
	todo := NewTodo(t)

	err := r.DB.Create(&todo).Error
	if err != nil {
		return err
	}

	return nil
}

// get all todo
func (r *Repository) ListTodos() ([]*entity.Todo, error) {
	var todos []entity.Todo

	err := r.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Todo, 0, len(todos))
	for _, todo := range todos {
		result = append(result, NewTodo(&todo))
	}

	return result, nil
}

// get todo by id
func (r *Repository) GetTodoByID(id int) (*entity.Todo, error) {
	var todo entity.Todo

	err := r.DB.First(&todo, id).Error
	if err != nil {
		return nil, err
	}

	return NewTodo(&todo), nil
}

// update todo information
func (r *Repository) PatchTodo(id int, t *entity.Todo) error {
	var todo *entity.Todo
	err := r.DB.First(&todo, id).Error
	if err != nil {
		return err
	}
	err = r.DB.Model(&todo).Updates(t).Error
	r.DB.Save(&todo)
	if err != nil {
		return err
	}

	return nil
}

// delete user by id
func (r *Repository) DeleteTodo(todo *entity.Todo) error {

	err := r.DB.First(&todo, todo.ID).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&todo).Error
	if err != nil {
		return err
	}

	return nil
}
