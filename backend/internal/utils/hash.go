package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256OfString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
