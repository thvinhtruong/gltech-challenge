package entity

type TodoUser struct {
	ID       int  `db: id`
	UserID   int  `db:"user_id"`
	TodoID   int  `db:"todo_id"`
	Finished bool `db:"is_finished"`
}
