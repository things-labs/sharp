package trans

import (
	"context"
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
