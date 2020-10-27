// Package trans gorm 事务封装
package trans

import (
	"context"

	"gorm.io/gorm"
)

// Trans 事务管理
type Trans struct {
	*gorm.DB
}

// New 创建事务管理实例
func New(db *gorm.DB) *Trans {
	return &Trans{db}
}

// Begin 开启事务,返回事务句柄
func (sf *Trans) Begin() *Trans {
	return &Trans{sf.DB.Begin()}
}

// ExecTrans 执行事务
// Deprecated: use Exec instead
func ExecTrans(ctx context.Context, db *gorm.DB, cb func(context.Context) error) error {
	return Exec(ctx, db, cb)
}

// Exec 执行事务
func Exec(ctx context.Context, db *gorm.DB, cb func(context.Context) error) error {
	if trans := FromTransCtx(ctx); trans != nil {
		return cb(ctx)
	}

	trans := New(db)
	tx := trans.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	ctx = NewTransCtx(ctx, tx)

	if err := cb(ctx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
