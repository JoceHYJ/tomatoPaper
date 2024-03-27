package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptionMd5(s string) string {
	pwd := md5.New()
	pwd.Write([]byte(s))
	return hex.EncodeToString(pwd.Sum(nil))
}
