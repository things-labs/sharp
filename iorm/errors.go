package iorm

import (
	"errors"
)

var (
	ErrZeroOrEmptyValue   = errors.New("value must not be zero or empty")
	ErrObjectAlreadyExist = errors.New("object already exist")
	ErrHasAssociation     = errors.New("object has associate other object")
)
