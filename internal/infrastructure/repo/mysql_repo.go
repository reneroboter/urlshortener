package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/reneroboter/urlshortener/internal/domain"
	"github.com/reneroboter/urlshortener/pkg/mysql"
)

type MySqlRepository struct {
	m *sql.DB
}

func NewMySqlStore() *MySqlRepository {
	return &MySqlRepository{
		m: mysql.NewMySqlClient(),
	}
}

func (s *MySqlRepository) Put(code, url string) error {
	result, err := s.m.Exec("INSERT INTO short_url (url, code) VALUES (?, ?)", url, code)
	if err != nil {
		return fmt.Errorf("ShortURL: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("ShortURL: %v", err)
	}
	slog.Info("LastInsertId: %d", id)
	return nil
}

func (s *MySqlRepository) Get(code string) (string, error) {
	var short domain.ShortURL

	row := s.m.QueryRow("SELECT url, code FROM short_url WHERE code = ?", code)
	if err := row.Scan(&short.URL, &short.Code); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("code not found")
		}
		return "", fmt.Errorf("ShortURL %d: %v", code, err)
	}

	return short.URL, nil
}
