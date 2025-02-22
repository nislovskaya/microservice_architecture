package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/httperrors"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/model"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}
type auth struct {
	logger    *logrus.Entry
	repo      repository.Repository
	secretKey string
}

func New(opts ...Option) Service {
	service := &auth{}

	for _, option := range opts {
		option(service)
	}

	return service
}

func (a *auth) Register(email, password string) error {
	exists, err := a.repo.ExistsByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if exists {
		return &httperrors.BadRequestError{
			Message: fmt.Sprintf("user with email %s already exists", email),
		}
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err = a.repo.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *auth) Login(email, password string) (string, error) {
	user, err := a.repo.GetByEmail(email)
	if err != nil {
		return "", &httperrors.UnauthorizedError{
			Message: "Invalid credentials",
			Err:     fmt.Errorf("failed to get user by email: %w", err),
		}
	}

	err = verifyPassword(user.Password, password)
	if err != nil {
		return "", &httperrors.UnauthorizedError{
			Message: "Invalid credentials",
			Err:     fmt.Errorf("failed to compare password: %w", err),
		}
	}

	token, err := a.generateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (a *auth) generateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
