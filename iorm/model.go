package iorm

import (
	"context"

	"github.com/jinzhu/gorm"
	icontext "github.com/thinkgos/assist/iorm/context"
)

// M 别名
type M map[string]interface{}

// GetDB get db
func GetDB(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if trans := icontext.FromTrans(ctx); trans != nil {
		if tx, ok := trans.(*gorm.DB); ok {
			return tx
		}
	}
	return defaultDB
}

// GetDBWithModel get db with model
func GetDBWithModel(ctx context.Context, defaultDB *gorm.DB, model interface{}) *gorm.DB {
	return GetDB(ctx, defaultDB).Model(model)
}
