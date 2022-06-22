package gormdb

import entity "github.com/thvinhtruong/legoha/app/domain/entities"

func newUser(user entity.User) entity.User {
	u := entity.User{
		ID:          user.ID,
		Name:        user.Name,
		Username:    user.Username,
		Password:    user.Password,
		Role:        user.Role,
		DateCreated: user.DateCreated,
	}
	return u
}

func newTodo(todo entity.Todo) entity.Todo {
	t := entity.Todo{}
	t.ID = todo.ID
	t.Title = todo.Title
	t.Description = todo.Description
	t.Completed = todo.Completed
	t.CreatedAt = todo.CreatedAt
	return t
}

func setUserRole(u entity.User, role string) {
	u.Role = "user"
}

func setTodoStatus(todo entity.TodoUser, value bool) {
	todo.Finished = value
}
