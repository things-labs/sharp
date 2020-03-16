package context

import (
	"context"
)

type transCtx struct{}
type transLockCtx struct{}

// NewTrans 创建事务的上下文
func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// FromTrans 从上下文中获取事务
func FromTrans(ctx context.Context) interface{} {
	return ctx.Value(transCtx{})
}

// NewTransLock 创建事务锁的上下文
func NewTransLock(ctx context.Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, struct{}{})
}

// FromTransLock 从上下文中获取事务锁
func FromTransLock(ctx context.Context) bool {
	return ctx.Value(transLockCtx{}) != nil
}
