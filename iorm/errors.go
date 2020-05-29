package iorm

import (
	"errors"
)

// 基本错误
var (
	ErrZeroOrEmptyValue    = errors.New("value must not be zero or empty")
	ErrObjectAlreadyExist  = errors.New("object already exist")
	ErrHasAssociation      = errors.New("object has associate other object")
	ErrOperateNotPermitted = errors.New("operate not permitted")
	ErrOutOfRange          = errors.New("out of range")
	ErrInvalidParameter    = errors.New("invalid parameter")
	ErrConflict            = errors.New("data field conflict")
)
