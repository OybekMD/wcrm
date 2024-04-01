package etc

import (
	"math/rand"
	"time"
)

func GenerateCode(length int) string {
	const charset = "0123456789"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}
