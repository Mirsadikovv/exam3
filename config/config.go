package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
	ServisName       string

	SmtpPassword string
	SmtpPort     string
	SmtpServer   string
	SmtpUsername string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("error!!!", err)
	}
	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "exam"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "mirodil"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1212"))
	cfg.ServisName = "LMS"

	cfg.SmtpPassword = cast.ToString(getOrReturnDefault("SMTP_PASSWORD", "mwzz ekyq ybqc wuzd"))
	cfg.SmtpPort = cast.ToString(getOrReturnDefault("SMTP_PORT", "587"))
	cfg.SmtpServer = cast.ToString(getOrReturnDefault("SMTP_SERVER", "smtp.gmail.com"))
	cfg.SmtpUsername = cast.ToString(getOrReturnDefault("SMTP_USERNAME", "mirsadikovmirodil52@gmail.com"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
