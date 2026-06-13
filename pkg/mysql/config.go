package mysql

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlConfig() string {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("MYSQL_ADDR")
	cfg.DBName = os.Getenv("MYSQL_DATABASE")

	return cfg.FormatDSN()
}
