package application

import (
	"log/slog"

	"github.com/reneroboter/urlshortener/internal/domain"
	"github.com/reneroboter/urlshortener/internal/infrastructure"
)

func NewShortURLService() ShortURLService {
	return ShortURLService{
		repo: *infrastructure.NewShortUrlRepository(),
	}
}

type ShortURLService struct {
	repo infrastructure.ShortURLRepository
}

func (s *ShortURLService) CreateShortURL(URL string) domain.ShortURL {
	shortURL := domain.ShortURL{
		URL:  URL,
		Code: HashUrl(NormalizeUrl(URL)),
	}

	if err := s.repo.Put(shortURL.Code, shortURL.URL); err != nil {
		slog.Error(err.Error())
	}

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
