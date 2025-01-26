package user

import (
	"fmt"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/model"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type user struct {
	logger *logrus.Entry
	repo   repository.Repository
}

func New(opts ...Option) Service {
	service := &user{}

	for _, option := range opts {
		option(service)
	}

	return service
}

func (s *user) CreateUser(user *model.User) error {
	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *user) GetUserByID(id uint) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *user) UpdateUser(user *model.User) error {
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *user) DeleteUser(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
