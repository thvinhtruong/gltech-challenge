package mysqldb

import (
	"context"
	"errors"

	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func (r *BaseRepository) CreateNewTodo(ctx context.Context, t entity.Todo) error {
	_, err := insertTodo(r, ctx, t)
	if err != nil {
		return err
	}
	return nil
}

// get todos according to flag: [0 - title], [1 - description], [2 - completed], [3 - all]
func (r *BaseRepository) GetTodosByFlags(ctx context.Context, title string, description string, flags int) ([]entity.Todo, error) {
	switch flags {
	case 0:
		return queryTodo(r, ctx, title, description, 0)
	case 1:
		return queryTodo(r, ctx, title, description, 1)
	case 2:
		return queryTodo(r, ctx, title, description, 2)
	}
	return []entity.Todo{}, errors.New("flag must be 0, 1, or 2")
}

func (r *BaseRepository) GetTodoByID(ctx context.Context, id int) (entity.Todo, error) {
	results, err := r.DB.QueryContext(ctx, sql_query_todo_id, id)
	if err != nil {
		return entity.Todo{}, err
	}

	var todo entity.Todo
	for results.Next() {
		err := results.Scan(&todo.Title, &todo.Description, &todo.CreatedAt)
		if err != nil {
			return entity.Todo{}, err
		}
	}

	return todo, nil
}

// update todo information
func (r *BaseRepository) PatchTodo(ctx context.Context, id int, t entity.Todo) error {
	stmt, err := r.DB.PrepareContext(ctx, sql_update_todo)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, t.Title, t.Description, id)
	if err != nil {
		return err
	}

	return nil
}

// delete user by id
func (r *BaseRepository) DeleteTodo(ctx context.Context, id int) error {
	stmt, err := r.DB.PrepareContext(ctx, sql_delete_todo)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
