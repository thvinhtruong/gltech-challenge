package registry

import (
	"database/sql"

	"github.com/thvinhtruong/legoha/app/interface/persistence/concrete"
	"github.com/thvinhtruong/legoha/app/interface/persistence/rdbms/mysqldb"
)

func GetRepository(db *sql.DB) mysqldb.BaseRepository {
	conn := concrete.DBConn{DB: db}
	return mysqldb.BaseRepository{DB: &conn}
}
