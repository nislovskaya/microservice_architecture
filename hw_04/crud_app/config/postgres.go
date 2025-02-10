package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(logger *logrus.Entry) *gorm.DB {
	host, err := GetConfigValue("DB_HOST")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	user, err := GetConfigValue("DB_USER")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	password, err := GetConfigValue("DB_PASSWORD")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	dbname, err := GetConfigValue("DB_NAME")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	logger.Infof("Database URL: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalln(fmt.Errorf("failed to connect database: %w", err))
	}
	return db
}
