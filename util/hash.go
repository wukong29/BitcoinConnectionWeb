package util

import (
	"crypto/md5"
	"encoding/hex"
)

//对一个字符串进行MD5哈希计算，并返回hash值
func MD5HashString(data string) string {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(data))
	return hex.EncodeToString(hashMd5.Sum(nil))
}
