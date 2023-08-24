package utils

import (
	"fmt"
	"testing"
)

// GenerateRandSecret 生成随机密钥 用于 jwt
func TestGenerateRandSecret(t *testing.T) {
	key := GenerateRandSecret(64)
	fmt.Println(key)
}
