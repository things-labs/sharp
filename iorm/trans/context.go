package trans

import (
	"context"

	"gorm.io/gorm"
)

type transCtx struct{}

// NewTransCtx 创建事务的上下文
func NewTransCtx(ctx context.Context, trans *Trans) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// FromTransCtx 从上下文中获取事务
func FromTransCtx(ctx context.Context) *Trans {
	value := ctx.Value(transCtx{})
	if value != nil {
		if trans, ok := value.(*Trans); ok {
			return trans
		}
	}
	return nil
}

// CtxDB ctx db 如果上下文中有事务,返回事务,否则使用db
func CtxDB(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tran := FromTransCtx(ctx); tran != nil {
			return tran.DB
		}
		return db
	}
}
