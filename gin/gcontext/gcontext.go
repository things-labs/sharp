package gcontext

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Context gin context
func Context(c *gin.Context) context.Context {
	return c.Request.Context()
}

// Value gin context value
func Value(c *gin.Context, key interface{}) interface{} {
	return Context(c).Value(key)
}
