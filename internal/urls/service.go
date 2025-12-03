package urls

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type Service struct {
	Repository *Repository
	BaseURL    string
}

func NewService(
	repo *Repository,
	baseURL string,
) *Service {
	return &Service{
		Repository: repo,
		BaseURL:    baseURL,
	}
}

const (
	codeLength  = 7
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateCode(n int) (string, error) {
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

func (s *Service) ShortenURL(ctx context.Context, longURL string) (string, error) {
	code, err := generateCode(codeLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate code: %w", err)
	}
	err = s.Repository.InsertURL(
		ctx,
		code,
		longURL,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate code: %w", err)
	}
	return fmt.Sprintf("%s/%s", s.BaseURL, code), nil
}

func (s *Service) GetOriginalURL(ctx context.Context, code string) (string, error) {
	longURL, err := s.Repository.GetURL(ctx, code)
	if err != nil {
		return "", fmt.Errorf("failed to get long URL: %w", err)
	}
	return longURL, nil
}
