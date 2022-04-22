package server

import (
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")
	return db
}
