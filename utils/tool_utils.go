package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/****
md5 返回32j机密小写
****/

func Md5String32(s string) string {
	h := md5.New()
	h.Write([]byte(s)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
