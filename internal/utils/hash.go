package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(password string, salt string) string {
	hash := md5.Sum([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}
