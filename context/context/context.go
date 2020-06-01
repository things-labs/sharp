package context

import (
	"context"
)

// traceID/requestid: github.com/thinkgos/gin-middlewares/tree/master/requestid

// 定义context 上下文中的键
type userNameCtx struct{}
type userIDCtx struct{}

// NewUserName 创建用户当前登录用户名的上下文
func NewUserName(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameCtx{}, userName)
}

// FromUserName 从上下文中获取用户名
func FromUserName(ctx context.Context) (string, error) {
	v := ctx.Value(userNameCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s, nil
		}
	}
	return "", ErrValueNotFound
}

// NewUserID 创建用户当前登录ID的上下文
func NewUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDCtx{}, userID)
}

// FromUserID 从上下文中获取用户ID
func FromUserID(ctx context.Context) (uint, error) {
	v := ctx.Value(userIDCtx{})
	if v != nil {
		if s, ok := v.(uint); ok {
			return s, nil
		}
	}
	return 0, ErrValueNotFound
}
