package utils

import "os"

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func HasEnv(key string) bool {
	value := os.Getenv(key)
	return value != ""
}
