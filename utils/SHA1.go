package utils

import (
	"crypto/sha1"
	"fmt"
)

// hash plaintext with SHA-1 对字符串加密
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
