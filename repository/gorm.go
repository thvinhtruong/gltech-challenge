package repository

import (
	"github.com/jinzhu/gorm"
	entity "github.com/thvinhtruong/legoha/entities"
)

type Repository struct {
	DB *gorm.DB
}

func (repo *Repository) Initialize(db *gorm.DB) *gorm.DB {
	repo.DB = db
	db.DropTableIfExists(&entity.User{})
	db.AutoMigrate(&entity.User{})

	db.DropTableIfExists(&entity.Todo{})
	db.AutoMigrate(&entity.Todo{})

	db.DropTableIfExists(&entity.TaskList{})
	db.AutoMigrate(&entity.TaskList{})
	return repo.DB
}
