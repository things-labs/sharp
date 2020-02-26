package password

import "golang.org/x/crypto/bcrypt"

var privateKey = "@#$%^&*()"

// RegisterPrivateKey 注册一个新的私有key,用于加盐
func RegisterPrivateKey(tk string) {
	privateKey = tk
}

// PrivateKey 获取私有key
func PrivateKey() string {
	return privateKey
}

// GenerateFromPassword 密码hash运算
func GenerateFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(privateKey+password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// CompareHashAndPassword 密码hash验证
func CompareHashAndPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(privateKey+password))
}
