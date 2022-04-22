package entity

import "github.com/jinzhu/gorm"

type TaskList struct {
	gorm.Model
	UserID   int  `db:"user_id"`
	TodoID   int  `db:"todo_id"`
	Finished bool `db:"is_finished"`
}
