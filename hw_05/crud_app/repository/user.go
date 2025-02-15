package repository

import (
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
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

func (u *user) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) Update(user *model.User) error {
	return u.db.Save(user).Error
}

func (u *user) Delete(id uint) error {
	return u.db.Delete(&model.User{}, id).Error
}
