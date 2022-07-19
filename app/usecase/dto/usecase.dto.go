package dto

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUser struct {
	TodoID int `json:"todo_id"`
	UserID int `json:"user_id"`
}
