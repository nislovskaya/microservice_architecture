package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Option func(r *routing)

func WithLogger(logger *logrus.Entry) Option {
	return func(r *routing) {
		r.logger = logger
	}
}

func WithDB(db *gorm.DB) Option {
	return func(r *routing) {
		r.db = db
	}
}
