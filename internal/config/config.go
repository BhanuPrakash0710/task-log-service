package config

import (
	"os"
)

func GetDBConfig() (string, string) {
	return os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
}
