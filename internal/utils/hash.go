package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data []byte) string {
	md5Instance := md5.New()
	md5Instance.Write(data)
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data[:])
}
