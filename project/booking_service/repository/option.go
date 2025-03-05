package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Option func(b *booking)

func WithLogger(logger *logrus.Entry) Option {
	return func(b *booking) {
		b.logger = logger
	}
}

func WithDB(db *gorm.DB) Option {
	return func(b *booking) {
		b.db = db
	}
}
