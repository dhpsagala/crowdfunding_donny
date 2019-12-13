package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashString(val string, key string) string {
	hasher := md5.New()
	hasher.Write([]byte(val))
	return hex.EncodeToString(hasher.Sum([]byte(key)))
}
