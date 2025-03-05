package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/handler"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/repository"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service"
	"github.com/nislovskaya/microservice_architecture/hw_06/auth_service/service/auth"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"net/http"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	router := getRouter()

	logger.Info("Server is started...")
	logger.Fatal(http.ListenAndServe(":8080", router))
}

func getRouter() *mux.Router {
	deps := initializeDependencies()
	services := initializeServices(deps)
	handlers := initializeHandlers(services)

	return handlers.InitRouter()
}

type dependencies struct {
	postgres      *gorm.DB
	secretKey     string
	redisClient   *redis.Client
	kafkaProducer *kafka.Producer
}

func initializeDependencies() *dependencies {
	postgres, err := config.ConnectPostgres(logger)
	if err != nil {
		logger.Fatal(err)
	}

	secretKey, err := config.GetSecret()
	if err != nil {
		logger.Fatal(err)
	}

	redisClient, err := config.ConnectRedis(logger)
	if err != nil {
		logger.Fatal(err)
	}

	kafkaProducer, err := kafka.NewProducer("kafka:9092")
	if err != nil {
		logger.Fatalf("Failed to create Kafka producer: %v", err)
	}

	return &dependencies{
		postgres:      postgres,
		secretKey:     secretKey,
		redisClient:   redisClient,
		kafkaProducer: kafkaProducer,
	}
}

func initializeServices(deps *dependencies) *service.Service {
	repo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(deps.postgres),
	)

	authService := auth.New(
		auth.WithLogger(logger),
		auth.WithRepository(repo),
		auth.WithSecretKey(deps.secretKey),
		auth.WithKafkaProducer(deps.kafkaProducer),
		auth.WithRedis(deps.redisClient),
	)

	return service.New(
		service.WithAuthService(authService),
	)
}

func initializeHandlers(services *service.Service) *handler.Handler {
	return handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)
}
