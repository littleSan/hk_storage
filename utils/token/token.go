package token

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GetToken() (str string) {
	str, _ = CreateToken(20)
	return
}
