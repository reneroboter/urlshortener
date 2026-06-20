package pkg

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Port int
	}
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
	}
	Redis struct {
		Host     string
		Port     int
		User     string
		Password string
	}
	Logging struct {
		Level string
	}
}

func Load() *Config {
	cfg := &Config{}

	cfg.Server.Port = castToInt(os.Getenv("SERVER_PORT"))

	cfg.MySQL.Host = os.Getenv("MYSQL_HOST")
	cfg.MySQL.Port = castToInt(os.Getenv("MYSQL_PORT"))
	cfg.MySQL.User = os.Getenv("MYQSL_USER")
	cfg.MySQL.User = os.Getenv("MYSQL_PASSWORD")

	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = castToInt(os.Getenv("REDIS_PORT"))
	cfg.Redis.User = os.Getenv("REDIS_USER")
	cfg.Redis.User = os.Getenv("REDIS_PASSWORD")

	cfg.Logging.Level = os.Getenv("MYSQL_PASSWORD")

	return cfg
}

func castToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		slog.Error("Can not cast parameter %s to integer", s)
	}
	return i
}
