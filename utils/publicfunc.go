package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(md5str string)string{
	h := md5.New()
	h.Write([]byte(md5str))
	return hex.EncodeToString(h.Sum(nil))
}