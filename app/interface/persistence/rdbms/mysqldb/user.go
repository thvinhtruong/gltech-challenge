package mysqldb

import (
	"context"
	"errors"
	"time"

	entity "github.com/thvinhtruong/legoha/app/domain/entities"
	"github.com/thvinhtruong/legoha/pkg/conversion"
)

// add a user
func (r *BaseRepository) CreateUser(ctx context.Context, user entity.User) error {
	_, err := insertUser(r, ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// get user by flags
func (r *BaseRepository) GetUserByFlags(ctx context.Context, name string, username string, flag int) ([]entity.User, error) {
	switch flag {
	case 0:
		return queryUser(r, ctx, name, username, 0)

	case 1:
		return queryUser(r, ctx, name, username, 1)

	case 2:
		return queryUser(r, ctx, name, username, 2)
	}

	return []entity.User{}, errors.New("invalid flag")
}

// get user by id
func (r *BaseRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	result, err := r.DB.QueryContext(ctx, sql_query_userid, id)
	if err != nil {
		return entity.User{}, err
	}

	var user entity.User
	for result.Next() {
		err := result.Scan(&user.Name, &user.Username, &user.Password, &user.CreatedAt)
		if err != nil {
			return entity.User{}, err
		}
	}

	return user, nil
}

// update user infor
func (r *BaseRepository) PatchUser(ctx context.Context, id int, u entity.User) error {
	stmt, err := r.DB.PrepareContext(ctx, sql_update_user)
	if err != nil {
		return err
	}
	defer stmt.Close()

	updated_at := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	_, err = stmt.ExecContext(ctx, u.Name, u.Username, u.Password, updated_at, id)
	if err != nil {
		return err
	}

	return nil
}

// delete user by id
func (r *BaseRepository) DeleteUser(ctx context.Context, id int) error {
	stmt, err := r.DB.PrepareContext(ctx, sql_delete_user)
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
