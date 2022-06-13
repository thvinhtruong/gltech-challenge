package entity

import "time"

type User struct {
	ID          int
	Name        string
	Username    string
	Password    string
	Role        string
	DateCreated time.Time
}
