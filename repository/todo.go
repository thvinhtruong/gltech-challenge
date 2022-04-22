package repository

import (
	entity "github.com/thvinhtruong/legoha/entities"
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

func (r *Repository) IsTodoExists(id int) bool {
	_, err := r.GetTodoByID(id)
	return err == nil
}

// add a todo
func (r *Repository) CreateNewTodo(entityTodo *entity.Todo) error {
	t := NewTodo(entityTodo)

	err := r.DB.Create(&t).Error
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
func (r *Repository) PatchTodo(id int, entityTodo *entity.Todo) error {
	t := NewTodo(entityTodo)
	err := r.DB.Model(&t).Where("id = ?", id).Update(t).Error
	if err != nil {
		return err
	}

	return nil
}

// delete user by id
func (r *Repository) DeleteTodo(id int) error {
	var todo entity.Todo

	err := r.DB.First(&todo, id).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&todo).Error
	if err != nil {
		return err
	}

	return nil
}
