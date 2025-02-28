package validation

import (
	"github.com/sirupsen/logrus"
)

type Option func(v *validation)

func WithLogger(logger *logrus.Entry) Option {
	return func(v *validation) {
		v.logger = logger
	}
}

func WithSecretKey(secretKey string) Option {
	return func(v *validation) {
		v.secretKey = secretKey
	}
}
