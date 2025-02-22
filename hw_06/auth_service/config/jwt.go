package config

import (
	"fmt"
)

func GetSecret() string {
	secretKey, err := GetConfigValue("JWT_SECRET")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	return secretKey
}
