package entity

type TodoUser struct {
	ID       int
	UserID   int  `db:"user_id"`
	TodoID   int  `db:"todo_id"`
	Finished bool `db:"is_finished"`
}
