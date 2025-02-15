package config

import (
	"os"
	"strings"
)

func GetConfigValue(key string) (string, error) {
	if filename := os.Getenv(key + "_FILE"); filename != "" {
		if data, err := os.ReadFile(filename); err == nil {
			return strings.Trim(string(data), "\n"), nil
		}
	}

	return os.Getenv(key), nil
}
