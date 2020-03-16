package iorm

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/thinkgos/assist/pagination"
)

var (
	ErrZeroOrEmptyValue = errors.New("value must not be zero or empty!")
)

// Query 查询所有
func Query(db *gorm.DB, out interface{}) error {
	return db.Find(out).Error
}

// QueryPage 分页查询,db需提供model和条件, list需提供切片地址 如 &[]yourStruct{}
// pg 如果均为默认参数,将不进行分页查询,将返回所有数据
func QueryPage(db *gorm.DB, pg pagination.Param, out interface{}) (pagination.Infos, error) {
	var total, pageIndex, pageSize int

	err := db.Count(&total).Error
	if err != nil {
		return pagination.Infos{}, err
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

	return pagination.Infos{
		Total:     total,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		List:      out,
	}, err
}

// QueryPageRelated 分页关联查询
//db需提供model(包含主键)和条件, list需提供切片地址 如 &[]yourStruct{}
// pg 如果均为默认参数,将不进行分页查询,将返回所有数据
func QueryPageRelated(db *gorm.DB, pg pagination.Param,
	out interface{}, foreignKeys ...string) error {
	if pg.PageSize > 0 {
		db = db.Limit(pg.PageSize)
		if pg.PageIndex > 0 {
			db = db.Offset(pg.PageSize * (pg.PageIndex - 1))
		}
	}
	return db.Related(out, foreignKeys...).Error
}

// QueryOne 根据id更新相应字段
func QueryOne(db *gorm.DB, query M, out interface{}) error {
	if len(query) == 0 {
		return db.First(out).Error
	}
	return db.First(out, query).Error
}

// UpdateStatus 根据id更新相应字段, db需提供model
func Update(db *gorm.DB, id uint, value M) error {
	if id == 0 || value == nil || len(value) == 0 {
		return ErrZeroOrEmptyValue
	}
	return db.Where("id=?", id).Update(value).Error
}
