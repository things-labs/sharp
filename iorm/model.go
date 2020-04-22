package iorm

import (
	"context"

	"github.com/jinzhu/gorm"
	icontext "github.com/thinkgos/assist/iorm/context"
)

// GetDB get db
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	if trans := icontext.FromTrans(ctx); trans != nil {
		if tx, ok := trans.(*gorm.DB); ok {
			return tx
		}
	}
	return defDB
}

// GetDBWithModel get db with model
func GetDBWithModel(ctx context.Context, defDB *gorm.DB, model interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(model)
}