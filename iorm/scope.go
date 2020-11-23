package iorm

import (
	"gorm.io/gorm"

	"github.com/thinkgos/sharp/core/paginator"
)

// Paginate paginate
func Paginate(pg paginator.Param) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pg.PageSize > 0 {
			db = db.Limit(pg.PageSize)
			if pg.PageIndex > 0 {
				db = db.Offset(pg.PageSize * (pg.PageIndex - 1))
			}
		}
		return db
	}
}
