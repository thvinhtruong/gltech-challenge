package presenter

import "errors"

var (
	Date = [...]string{"Mon", "Tues", "Weds", "Thu", "Fri", "Sat", "Sun"}
)

func ValidateAuth(username string, password string) error {
	if len(username) == 0 || len(password) == 0 {
		return errors.New("username or password is empty")
	}

	return nil
}

func ValidateDate(str string) error {
	for _, v := range Date {
		if str == v {
			return nil
		}
	}

	return errors.New("invalid date")
}

func ValidateTodo(title string) error {
	if len(title) == 0 {
		return errors.New("title is empty")
	}

	return nil
}
