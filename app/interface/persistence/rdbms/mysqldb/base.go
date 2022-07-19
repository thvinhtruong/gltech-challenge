package mysqldb

import (
	"github.com/thvinhtruong/legoha/app/interface/persistence/concrete"
)

type BaseRepository struct {
	DB concrete.DBTX
}

// EnableTx allow the transaction to be initialized.
func (q *BaseRepository) EnableTx(txFunc func() error) error {
	tx, err := q.DB.TxBegin()
	if err != nil {
		return err
	}

	return tx.TxEnd(txFunc)
}

func NewBaseRepository(DB concrete.DBTX) *BaseRepository {
	return &BaseRepository{DB}
}
