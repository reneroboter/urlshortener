package mysql

import (
	"database/sql"
	"log/slog"
)

func NewMySqlClient() *sql.DB {
	var err error
	db, err := sql.Open("mysql", NewMySqlConfig())
	if err != nil {
		slog.Error(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		slog.Error(pingErr.Error())
	}
	slog.Info("MySQL is Connected!")
	return db
}
