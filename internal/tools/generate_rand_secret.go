package tools

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandSecret 生成随机密钥 用于 jwt
func GenerateRandSecret(length int) string {
	key := make([]byte, length/4*3)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
