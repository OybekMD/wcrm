package etc

import (
	"crypto/rand"
	"fmt"
	"io"
)

// Table for code generator
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// Generate code is function that create n-digit random code
func GenerateCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		fmt.Println(err)
		return ""
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)

}

/*
func GenerateCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}
*/