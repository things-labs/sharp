package iorm

import (
	"gorm.io/gorm"

	"github.com/thinkgos/sharp/core/paginator"
)

// QueryPages 分页查询
// db需提供model和条件, list需提供切片地址 如 &[]yourStruct{}
// pg 如果均为默认参数,将不进行分页查询,将返回所有数据
func QueryPages(db *gorm.DB, pg paginator.Param, out interface{}) (paginator.Info, error) {
	var total int64
	var pageIndex, pageSize int

	err := db.Count(&total).Error
	if err != nil {
		return paginator.Info{}, err
	}
	if pg.PageSize > 0 {
		pageSize = pg.PageSize
		db = db.Limit(pageSize)
		if pg.PageIndex > 0 {
			pageIndex = pg.PageIndex
			db = db.Offset(pageSize * (pageIndex - 1))
		}
	}
	err = db.Find(out).Error
	return paginator.Info{
		Total:     total,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}, err
}

// QueryPagesAssociation query page association
func QueryPagesAssociation(db *gorm.DB, pg paginator.Param, out interface{}, column string) (paginator.Info, error) {
	var pageIndex, pageSize int

	total := db.Association(column).Count()
	if pg.PageSize > 0 {
		pageSize = pg.PageSize
		db = db.Limit(pageSize)
		if pg.PageIndex > 0 {
			pageIndex = pg.PageIndex
			db = db.Offset(pageSize * (pageIndex - 1))
		}
	}
	err := db.Association(column).Find(out)
	return paginator.Info{
		Total:     total,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}, err
}
