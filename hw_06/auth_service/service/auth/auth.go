package auth

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/httperrors"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/model"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	Register(email, password string) (uint, error)
	Login(email, password string) (string, error)
	ValidateToken(token string) (*model.JwtClaims, error)
	Logout(token string) error
}

type auth struct {
	logger    *logrus.Entry
	repo      repository.Repository
	secretKey string
	kafka     *kafka.Producer
	redis     *redis.Client
}

func New(opts ...Option) Service {
	service := &auth{}

	for _, option := range opts {
		option(service)
	}

	return service
}

func (a *auth) Register(email, password string) (uint, error) {
	exists, err := a.repo.ExistsByEmail(email)
	if err != nil {
		return 0, fmt.Errorf("failed to check if user exists: %w", err)
	}
	if exists {
		return 0, &httperrors.ConflictError{
			Message: fmt.Sprintf("user with email %s already exists", email),
		}
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err = a.repo.Create(user); err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	event := model.Event{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: time.Now(),
	}

	if err = a.kafka.Publish("user-events", event); err != nil {
		a.logger.Errorf("Failed to publish user created event: %v", err)
	}

	return user.ID, nil
}

func (a *auth) Login(email, password string) (string, error) {
	user, err := a.repo.GetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", &httperrors.UnauthorizedError{
			Message: "invalid credentials",
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.JwtClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func (a *auth) ValidateToken(tokenString string) (*model.JwtClaims, error) {
	blacklisted, err := a.redis.Exists(context.Background(), fmt.Sprintf("blacklist:%s", tokenString)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to check blacklist: %w", err)
	}
	if blacklisted > 0 {
		return nil, fmt.Errorf("token is blacklisted")
	}

	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func (a *auth) Logout(token string) error {
	claims, err := a.ValidateToken(token)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	exp := time.Unix(claims.ExpiresAt.Unix(), 0)
	ttl := time.Until(exp)

	if err := a.redis.Set(context.Background(), fmt.Sprintf("blacklist:%s", token), 1, ttl).Err(); err != nil {
		return fmt.Errorf("failed to blacklist token: %w", err)
	}

	return nil
}
