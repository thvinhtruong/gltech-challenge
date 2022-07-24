package sqlconn

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/thvinhtruong/legoha/app/config"
)

var (
	DB *sql.DB
)

func Init(cfg config.Config) error {
	DB, err := OpenConnection(cfg)
	if err != nil {
		return err
	}

	log.Println("Connected to database. MySQL version: ", DB.Ping())

	return nil
}

func OpenConnection(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&tls=%s&timeout=10s",
		cfg.DB_User,
		cfg.DB_Pass,
		cfg.DB_Host,
		cfg.DB_Name,
		cfg.DB_SSLMode,
	))

	if err != nil {
		log.Println("Exit with error: ", err)
		time.Sleep(time.Second * 3)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
