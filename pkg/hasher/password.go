package hasher

import (
	"crypto/sha256"
	"fmt"
)

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{
		salt: salt,
	}
}

func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
