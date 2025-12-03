package urls

import (
	"context"
	"fmt"

	"github.com/calqs/frogshort/pkg/code"
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
	codeLength = 7
)

func (s *Service) ShortenURL(ctx context.Context, longURL string) (string, error) {
	code, err := code.Generate(codeLength)
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
