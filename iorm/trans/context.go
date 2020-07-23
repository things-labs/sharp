package trans

import (
	"context"
)

type transCtx struct{}
type transLockCtx struct{}

// NewTransCtx 创建事务的上下文
func NewTransCtx(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// FromTransCtx 从上下文中获取事务
func FromTransCtx(ctx context.Context) interface{} {
	return ctx.Value(transCtx{})
}

// NewTransLockCtx 创建事务锁的上下文
func NewTransLockCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, struct{}{})
}

// FromTransLockCtx 从上下文中获取事务锁
func FromTransLockCtx(ctx context.Context) bool {
	return ctx.Value(transLockCtx{}) != nil
}
