package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/model"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error

	StartConsumer(ctx context.Context)
}

type user struct {
	logger   *logrus.Entry
	repo     repository.Repository
	consumer *kafka.Consumer
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

func (s *user) StartConsumer(ctx context.Context) {
	if s.consumer == nil {
		s.logger.Error("Consumer is not initialized")
		return
	}

	go func() {
		s.logger.Info("Starting Kafka consumer")
		if err := s.consumer.Consume(ctx, s.handleMessage); err != nil {
			s.logger.Errorf("Failed to consume message: %v", err)
		}
	}()
}

func (s *user) handleMessage(message []byte) error {
	var event model.UserCreatedEvent
	if err := json.Unmarshal(message, &event); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	s.logger.Infof("Received event for user with ID %d and email %s", event.ID, event.Email)

	// Создаем пользователя на основе события
	user := &model.User{
		ID:        event.ID,
		Email:     event.Email,
		FirstName: "",
		LastName:  "",
		Username:  event.Email, // Используем email как временное имя пользователя
		Phone:     "",
	}

	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("failed to create user from event: %w", err)
	}

	s.logger.Infof("Created user with ID %d from event", user.ID)
	return nil
}
