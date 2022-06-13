package gormdb

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// get the transaction to db from context
func (repo *Repository) getTx(ctx context.Context) *gorm.DB {
	v := ctx.Value("tx")
	if v == nil {
		panic(fmt.Sprintf("%s: no transaction found in context", "repository"))
	}
	tx, ok := v.(*gorm.DB)
	if !ok {
		panic(fmt.Sprintf("%s: can not get gorm DB in context: %T", "repository", v))
	}
	return tx
}
