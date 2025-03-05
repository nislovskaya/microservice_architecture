package repository

import (
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	ExistsByEmail(email string) (bool, error)
}

type user struct {
	logger *logrus.Entry
	db     *gorm.DB
}

func New(opts ...Option) Repository {
	repository := &user{}

	for _, option := range opts {
		option(repository)
	}

	return repository
}

func (u *user) Create(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *user) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := u.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
