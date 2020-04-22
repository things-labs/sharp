package password

import "golang.org/x/crypto/bcrypt"

type BCrypt struct {
	key string
}

func NewBCrypt(privateKey string) *BCrypt {
	if privateKey == "" {
		privateKey = defaultPrivateKey
	}
	return &BCrypt{privateKey}
}

// Hash 密码hash运算
func (sf BCrypt) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(sf.key+password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// Compare 密码hash验证
func (sf BCrypt) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(sf.key+password))
}
