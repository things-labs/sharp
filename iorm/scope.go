package iorm

import (
	"context"

	"gorm.io/gorm"

	"github.com/thinkgos/sharp/core/paginator"
	"github.com/thinkgos/sharp/iorm/trans"
)

// Paginate paginate
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

// CtxDB ctx db 如果上下文中有事务,返回事务,否则使用db
func CtxDB(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tran := trans.FromTransCtx(ctx); tran != nil {
			return tran.DB
		}
		return db
	}
}
