package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Option func(u *user)

func WithLogger(logger *logrus.Entry) Option {
	return func(u *user) {
		u.logger = logger
	}
}

func WithDB(db *gorm.DB) Option {
	return func(u *user) {
		u.db = db
	}
}
