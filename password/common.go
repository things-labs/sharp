package password

import (
	"errors"
)

const defaultPrivateKey = "@#$%^&*()"

var ErrCompareFailed = errors.New("verify compare failed")

type Verify interface {
	Hash(password string) (string, error)
	Compare(password, hash string) error
}
