package entity

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Completed   bool
}
