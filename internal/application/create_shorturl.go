package application

import (
	"log/slog"

	"github.com/reneroboter/urlshortener/internal/domain"
	"github.com/reneroboter/urlshortener/internal/infrastructure"
)

func NewShortURLService() ShortURLService {
	return ShortURLService{
		repo:      infrastructure.NewShortUrlRepository(),
		generator: infrastructure.SHA1CodeGenerator{},
	}
}

func NewTestShortURLService() ShortURLService {
	return ShortURLService{
		repo:      infrastructure.NewInMemoryStore(),
		generator: infrastructure.SHA1CodeGenerator{},
	}
}

type ShortURLService struct {
	repo      infrastructure.RepositoryInterface
	generator infrastructure.CodeGenerator
}

func (s *ShortURLService) CreateShortURL(URL string) domain.ShortURL {
	shortURL := domain.ShortURL{
		URL:  URL,
		Code: s.generator.GenerateCode(NormalizeUrl(URL)),
	}

	if err := s.repo.Put(shortURL.Code, shortURL.URL); err != nil {
		slog.Error(err.Error())
	}

	slog.Info(shortURL.URL)

	return shortURL
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
