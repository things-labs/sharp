// Package trans gorm 事务封装
package trans

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	icontext "github.com/thinkgos/sharp/iorm/context"
)

// Func 定义事务执行函数
type Func func(context.Context) error

// Trans 事务管理
type Trans struct {
	db *gorm.DB
}

// New 创建事务管理实例
func New(db *gorm.DB) *Trans {
	return &Trans{db}
}

// Begin 开启事务,返回事务句柄
func (a *Trans) Begin() (interface{}, error) {
	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// Commit 提交事务
func (a *Trans) Commit(trans interface{}) error {
	tx, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("unknown trans")
	}
	return tx.Commit().Error
}

// Rollback 回滚事务
func (a *Trans) Rollback(trans interface{}) error {
	tx, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("unknown trans")
	}
	return tx.Rollback().Error
}

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, db *gorm.DB, tf Func) error {
	if trans := icontext.FromTrans(ctx); trans != nil {
		return tf(ctx)
	}

	transModel := New(db)
	trans, err := transModel.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = transModel.Rollback(trans)
			panic(r)
		}
	}()

	ctx = icontext.NewTrans(ctx, trans)
	err = tf(ctx)
	if err != nil {
		_ = transModel.Rollback(trans)
		return err
	}
	return transModel.Commit(trans)
}

// ExecTransWithLock 执行事务（加锁）
func ExecTransWithLock(ctx context.Context, db *gorm.DB, cb Func) error {
	if !icontext.FromTransLock(ctx) {
		ctx = icontext.NewTransLock(ctx)
	}
	return ExecTrans(ctx, db, cb)
}
