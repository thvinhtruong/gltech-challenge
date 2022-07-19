package mysqldb

import (
	"github.com/thvinhtruong/legoha/app/interface/persistence/concrete"
	"github.com/thvinhtruong/legoha/pkg/testhelper"
)

// helper creates functions for unit test in data layer

func NewTestRepository(testName string) BaseRepository {
	db, err := testhelper.OpenConnectionForTest()
	if err != nil {
		panic(err)
	}

	// no transaction in data layer
	return BaseRepository{DB: &concrete.DBConn{DB: db}}
}
