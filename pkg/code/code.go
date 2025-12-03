package code

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Generate(n int) (string, error) {
	var sb strings.Builder
	sb.Grow(n)

	for range n {
		max := big.NewInt(int64(len(base62Chars)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		sb.WriteByte(base62Chars[num.Int64()])
	}

	return sb.String(), nil
}
