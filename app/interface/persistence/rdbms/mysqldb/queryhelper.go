package mysqldb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	"github.com/thvinhtruong/legoha/pkg/conversion"
)

// helper for complext query strings

// user
var (
	sql_insert_user    = `INSERT INTO users (name, username, password, role, created_at) VALUES (?, ?, ?, ?, ?)`
	sql_update_user    = `UPDATE users SET name = ?, username = ?, password = ?, updated_at = ? WHERE id = ?`
	sql_delete_user    = `DELETE FROM users WHERE id = ?`
	sql_query_name     = `SELECT name, username, created_at FROM users WHERE name LIKE ?`
	sql_query_username = `SELECT name, username, created_at FROM users WHERE username LIKE ?`
	sql_query_userid   = `SELECT name, username, created_at FROM users WHERE id = ?`
)

// todo
var (
	sql_insert_todo            = `INSERT INTO todo (title, description, created_at) VALUES (?, ?, ?)`
	sql_update_todo            = `UPDATE todo SET title = ?, description = ?, updated_at = ? WHERE id = ?`
	sql_delete_todo            = `DELETE FROM todo WHERE id = ?`
	sql_query_todo_id          = `SELECT title, description, created_at FROM todo WHERE id = ?`
	sql_query_todo_title       = `SELECT title, description, created_at FROM todo WHERE title LIKE ?`
	sql_query_todo_description = `SELECT title, description, created_at FROM todo WHERE description LIKE '% ? %'`
	sql_query_todo_created_at  = `SELECT title, description, created_at FROM todo WHERE created_at = ?`
)

// todo user
var (
	sql_assign       = `INSERT INTO todo_user (user_id, todo_id) VALUES (?, ?)`
	sql_revoke       = `DELETE FROM todo_user WHERE todo_id = ? AND user_id = ?`
	sql_query_todoid = `SELECT user_id FROM todo_user WHERE todo_id = ?`
	sql_status       = `UPDATE INTO todo_user SET finished = ? WHERE todo_id = ? AND user_id = ?`
)

// insert user into db
func insertUser(b *BaseRepository, ctx context.Context, user entity.User) (int, error) {
	role := "user"
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	results, err := b.DB.ExecContext(ctx, sql_insert_user, user.Name, user.Username, user.Password, role, date)
	if err != nil {
		return 0, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// read necessary fields of users when querrying
func readUserFields(results *sql.Rows) []entity.User {
	users := []entity.User{}
	for results.Next() {
		var user entity.User
		err := results.Scan(&user.Name, &user.Username, &user.CreatedAt)
		if err != nil {
			return []entity.User{}
		}
		users = append(users, user)
	}
	return users
}

// read user according to flag: [0 - fullname] [1 - username] [2 - all]
func queryUser(b *BaseRepository, ctx context.Context, name string, username string, flag int) ([]entity.User, error) {
	switch flag {
	case 0:
		if name == "" {
			return []entity.User{}, errors.New("name must not be empty")
		}

		results, err := b.DB.QueryContext(ctx, sql_query_name, username)
		if err != nil {
			return nil, err
		}

		return readUserFields(results), nil

	case 1:
		if username == "" {
			return []entity.User{}, errors.New("username must not be empty")
		}

		results, err := b.DB.QueryContext(ctx, sql_query_username, username)
		if err != nil {
			return nil, err
		}

		return readUserFields(results), nil

	case 2:
		results, err := b.DB.QueryContext(ctx, "select name, username, created_at from user")
		if err != nil {
			return nil, err
		}

		return readUserFields(results), nil
	}

	return []entity.User{}, errors.New("flag must be 0, 1 or 2")
}

func insertTodo(b *BaseRepository, ctx context.Context, todo entity.Todo) (int, error) {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	results, err := b.DB.ExecContext(ctx, sql_insert_todo, todo.Title, todo.Description, date)
	if err != nil {
		return 0, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func readTodoField(results *sql.Rows) []entity.Todo {
	todos := []entity.Todo{}
	for results.Next() {
		var todo entity.Todo
		err := results.Scan(&todo.Title, &todo.Description, &todo.CreatedAt)
		if err != nil {
			return []entity.Todo{}
		}
		todos = append(todos, todo)
	}
	return todos
}

// get todo according to flag: [0 - title], [1 - description], [2 - all]
func queryTodo(b *BaseRepository, ctx context.Context, title string, description string, flag int) ([]entity.Todo, error) {
	switch flag {
	case 0:
		if title == "" {
			return []entity.Todo{}, errors.New("title must not be empty")
		}

		results, err := b.DB.QueryContext(ctx, sql_query_todo_title, title)
		if err != nil {
			return nil, err
		}

		return readTodoField(results), nil

	case 1:
		// description can be empty
		results, err := b.DB.QueryContext(ctx, sql_query_todo_description, description)
		if err != nil {
			return nil, err
		}

		return readTodoField(results), nil

	case 2:
		results, err := b.DB.QueryContext(ctx, "select title, description, created_at from todo")
		if err != nil {
			return nil, err
		}

		return readTodoField(results), nil
	}

	return []entity.Todo{}, errors.New("flag must be 0, 1, 2 or 3")
}

func queryTodoOnDate(b *BaseRepository, ctx context.Context, created_at int64) ([]entity.Todo, error) {
	results, err := b.DB.QueryContext(ctx, sql_query_todo_created_at, created_at)
	if err != nil {
		return nil, err
	}

	return readTodoField(results), nil
}

func queryTodoFromDate(b *BaseRepository, ctx context.Context, date_begin int64, date_end int64) ([]entity.Todo, error) {
	sql_query_todo_date_range := "select title, description, created_at from todo where created_at between ? and ?"
	results, err := b.DB.QueryContext(ctx, sql_query_todo_date_range, date_begin, date_end)
	if err != nil {
		return nil, err
	}

	return readTodoField(results), nil
}

func queryUserFromTodoId(b *BaseRepository, ctx context.Context, todo_id int) ([]entity.User, error) {
	userid_rows, err := b.DB.QueryContext(ctx, sql_query_todoid, todo_id)
	if err != nil {
		return nil, err
	}

	userids := []int{}
	for userid_rows.Next() {
		var user_id int
		err := userid_rows.Scan(&user_id)
		if err != nil {
			return nil, err
		}
		userids = append(userids, user_id)
	}

	results := []entity.User{}
	for _, id := range userids {
		users, err := b.DB.QueryContext(ctx, sql_query_userid, id)
		if err != nil {
			return nil, err
		}
		results = append(results, readUserFields(users)...)
	}

	if len(results) == 0 {
		return nil, errors.New("no user found for this todo")
	}

	return results, nil
}

func queryTodoFromUserId(b *BaseRepository, ctx context.Context, user_id int) ([]entity.Todo, error) {
	todo_rows, err := b.DB.QueryContext(ctx, "select todo_id from todo_user where user_id = ?", user_id)
	if err != nil {
		return nil, err
	}

	todoids := []int{}
	for todo_rows.Next() {
		var todoid int
		err := todo_rows.Scan(&todoid)
		if err != nil {
			return nil, err
		}
		todoids = append(todoids, todoid)
	}

	results := []entity.Todo{}
	for _, id := range todoids {
		todos, err := b.DB.QueryContext(ctx, sql_query_todo_id, id)
		if err != nil {
			return nil, err
		}
		results = append(results, readTodoField(todos)...)
	}

	if len(results) == 0 {
		return nil, errors.New("no todo found for this user")
	}

	return results, nil
}

func assignHelper(b *BaseRepository, ctx context.Context, todouser entity.TodoUser) (int, error) {
	result, err := b.DB.ExecContext(ctx, sql_assign, todouser.UserID, todouser.TodoID)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows == 0 {
		return 0, errors.New("no user or todo found for this assignment")
	}

	ids, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(ids), nil
}

func revokeHelper(b *BaseRepository, ctx context.Context, todoId int, userId int) error {
	stmt, err := b.DB.PrepareContext(ctx, sql_revoke)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, userId, todoId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user or todo not found")
	}

	return nil
}

func isFinishedHelper(b *BaseRepository, ctx context.Context, finished bool, user_id int, todo_id int) error {
	stmt, err := b.DB.PrepareContext(ctx, sql_status)
	if err != nil {
		return err
	}
	status := conversion.ConvertBool2Int(finished)
	result, err := stmt.ExecContext(ctx, status, user_id, todo_id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user or todo not found")
	}

	return nil
}
