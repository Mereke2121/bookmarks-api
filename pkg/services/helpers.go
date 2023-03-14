package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

const (
	salt = "adfhlahfpdaohnvfoshj8943jf943jf8943hjriomjf8e3"
)

func generatePassword(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	return hex.EncodeToString(h.Sum([]byte(salt)))
}
