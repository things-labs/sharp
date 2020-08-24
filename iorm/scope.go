package iorm

import (
	"context"

	"gorm.io/gorm"

	"github.com/thinkgos/sharp/core/paginator"
	"github.com/thinkgos/sharp/iorm/trans"
)

func Paginate(pg paginator.Param) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pg.PageSize > 0 {
			db = db.Limit(pg.PageSize)
			if pg.PageIndex > 0 {
				db = db.Offset(pg.PageSize * (pg.PageIndex - 1))
			}
		}
		return db
	}
}

func CtxDB(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tran := trans.FromTransCtx(ctx); tran != nil {
			if tx, ok := tran.(*gorm.DB); ok {
				return tx
			}
		}
		return db
	}
}
