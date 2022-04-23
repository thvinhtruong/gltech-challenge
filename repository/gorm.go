package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB *gorm.DB
}
