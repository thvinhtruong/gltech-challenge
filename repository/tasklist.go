package repository

import entity "github.com/thvinhtruong/legoha/entities"

func (r *Repository) Assign(tl *entity.TaskList) error {
	if !r.IsUserExists(tl.UserID) && !r.IsTodoExists(tl.UserID) {
		return nil
	} else {
		t := entity.TaskList{UserID: tl.UserID, TodoID: tl.TodoID, Finished: false}
		err := r.DB.Create(&t).Error
		if err != nil {
			return err
		}

		return nil
	}
}

func (r *Repository) ListUsersForOneTodo(tl *entity.TaskList) ([]*entity.User, error) {
	var users []entity.User

	err := r.DB.Find(&users).Where("id = ?", tl.UserID).Error
	if err != nil {
		return nil, err
	}

	result := make([]*entity.User, 0, len(users))
	for _, user := range users {
		result = append(result, NewUser(&user))
	}
	return result, err

}

func (r *Repository) Completed(tl *entity.TaskList) error {
	err := r.DB.Find(&tl).Where("user_id = ? AND todo_id = ?", tl.UserID, tl.TodoID).Error
	if err != nil {
		return err
	}

	tl.Finished = true
	err = r.DB.Save(&tl).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Undo(tl *entity.TaskList) error {
	err := r.DB.Find(&tl).Where("user_id = ? AND todo_id = ?", tl.UserID, tl.TodoID).Error
	if err != nil {
		return err
	}

	tl.Finished = false
	err = r.DB.Save(&tl).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Revoke(tl entity.TaskList) error {
	err := r.DB.Delete(&tl).Error
	if err != nil {
		return err
	}

	return nil
}
