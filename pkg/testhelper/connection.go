package testhelper

import (
	"database/sql"
	"fmt"
)

const (
	User     = "root"
	Password = "root"
	Host     = "127.0.0.1:3306"
	Name     = "school"
)

func OpenConnectionForTest() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Name,
	))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
