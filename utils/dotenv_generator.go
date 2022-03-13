package utils

import (
	"os"
	"github.com/joho/godotenv"
)

func DotEnvGenerator(key string) (value string) {
	err := godotenv.Load(".env")
	CheckError(err)

	return os.Getenv(key)
}