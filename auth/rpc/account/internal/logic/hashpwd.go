package logic

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func pwd2Key(pwd string, salt string) (string, error) {
	key, err := scrypt.Key([]byte(pwd), []byte(salt), 1<<16, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}
