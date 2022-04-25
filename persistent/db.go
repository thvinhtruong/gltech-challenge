package persistent

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"github.com/thvinhtruong/legoha/config"
	entity "github.com/thvinhtruong/legoha/entities"
)

var params = config.LoadEnvironmentFile("config.env")
var db_host = params[0]
var db_port = params[1]
var db_user = params[2]
var db_pass = params[3]
var db_name = params[4]

func CreateDB() {
	credentials := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/"
	db, err := sql.Open("mysql", credentials)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + db_name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + db_name)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	CreateDB()
	credentials := db_user + ":" + db_pass + "@/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", credentials)
	if err != nil {
		panic(err)
	}

	db.DropTableIfExists(&entity.User{})
	db.AutoMigrate(&entity.User{})
	db.DropTableIfExists(&entity.Todo{})
	db.AutoMigrate(&entity.Todo{})
	db.DropTableIfExists(&entity.TaskList{})
	db.AutoMigrate(&entity.TaskList{})

	db.Exec("INSERT INTO users(Name, Username, Password, Role) VALUES('admin', 'admin', 'admin', 'admin')")

	return db
}
