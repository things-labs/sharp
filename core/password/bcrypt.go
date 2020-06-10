package password

import "golang.org/x/crypto/bcrypt"

// BCrypt bcrypt password encryption
type BCrypt struct {
	key string
}

// NewBCrypt new bcrypt password encryption with key
func NewBCrypt(privateKey string) *BCrypt {
	if privateKey == "" {
		privateKey = defaultPrivateKey
	}
	return &BCrypt{privateKey}
}

// Hash password hash encryption
func (sf BCrypt) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(sf.key+password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// Compare password hash verification
func (sf BCrypt) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(sf.key+password))
}
