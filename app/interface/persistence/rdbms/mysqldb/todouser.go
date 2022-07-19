package mysqldb

import (
	"context"

	entity "github.com/thvinhtruong/legoha/app/domain/entities"
)

func (r *BaseRepository) Assign(ctx context.Context, td entity.TodoUser) (int, error) {
	return assignHelper(r, ctx, td)
}

func (r *BaseRepository) ListUsersForOneTodo(ctx context.Context, todo_id int) ([]entity.User, error) {
	return queryUserFromTodoId(r, ctx, todo_id)
}

func (r *BaseRepository) ListTodosForOneUser(ctx context.Context, user_id int) ([]entity.Todo, error) {
	return queryTodoFromUserId(r, ctx, user_id)
}

func (r *BaseRepository) Completed(ctx context.Context, todoId int, userId int) error {
	return isFinishedHelper(r, ctx, true, todoId, userId)
}

func (r *BaseRepository) Undo(ctx context.Context, todoId int, userId int) error {
	return isFinishedHelper(r, ctx, false, todoId, userId)
}

func (r *BaseRepository) Revoke(ctx context.Context, todoId int, userId int) error {
	return revokeHelper(r, ctx, todoId, userId)
}
