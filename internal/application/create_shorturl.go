package application

import (
	"log/slog"

	"github.com/reneroboter/urlshortener/internal/domain"
	"github.com/reneroboter/urlshortener/internal/infrastructure"
	"github.com/reneroboter/urlshortener/internal/infrastructure/repo"
)

func NewShortURLService() ShortURLService {
	return ShortURLService{
		repo:      repo.NewShortUrlRepository(),
		generator: infrastructure.SHA1CodeGenerator{},
	}
}

func NewTestShortURLService() ShortURLService {
	return ShortURLService{
		repo:      repo.NewInMemoryStore(),
		generator: infrastructure.SHA1CodeGenerator{},
	}
}

type ShortURLService struct {
	repo      repo.RepositoryInterface
	generator infrastructure.CodeGenerator
}

func (s *ShortURLService) CreateShortURL(URL string) (domain.ShortURL, error) {
	shortURL := domain.ShortURL{
		URL:  URL,
		Code: s.generator.GenerateCode(NormalizeUrl(URL)),
	}

	if err := s.repo.Put(shortURL.Code, shortURL.URL); err != nil {
		slog.Error("failed to persist short URL",
			"code", shortURL.Code,
			"url", URL,
			"error", err,
		)
		return shortURL, err
	}

	return shortURL, nil
}

func (s *ShortURLService) ReturnShortURL(Code string) (error, domain.ShortURL) {
	URL, err := s.repo.Get(Code)
	if err != nil {
		return err, domain.ShortURL{}
	}

	return nil, domain.ShortURL{
		URL:  URL,
		Code: Code,
	}
}
