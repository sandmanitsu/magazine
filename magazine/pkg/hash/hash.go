package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type IHasher interface {
	Hash(pass string) (string, error)
	Compare(hashed, input string) bool
}

type BcryptHasher struct {
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) Hash(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hashedPassword), nil
}

func (h *BcryptHasher) Compare(hashed, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))

	return err == nil
}
