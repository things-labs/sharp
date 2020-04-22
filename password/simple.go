package password

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

type Simple struct {
	key string
}

func NewSimple(key string) *Simple {
	if key == "" {
		key = defaultPrivateKey
	}
	return &Simple{key}
}

// Hash 生成密钥 加盐法 md5Pwd+`@#$%`+md5Pwd+`^&*()`拼接
func (sf Simple) Hash(password string) (string, error) {
	h := md5.New()
	_, _ = io.WriteString(h, password+sf.key)

	md5Pwd := hex.EncodeToString(h.Sum(nil))
	// 加盐值加密
	_, _ = io.WriteString(h, sf.key)
	_, _ = io.WriteString(h, password)
	_, _ = io.WriteString(h, `@#$%`+sf.key)
	_, _ = io.WriteString(h, md5Pwd)
	_, _ = io.WriteString(h, `^&*()`+sf.key)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (sf Simple) Compare(password, hash string) error {
	h, err := sf.Hash(password)
	if err != nil {
		return err
	}
	if hash != h {
		return ErrCompareFailed
	}
	return nil
}
