package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/handler"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/repository"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/service"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/service/user"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	router := getRouter()

	logger.Info("Server is started...")
	logger.Fatal(http.ListenAndServe(":8081", router))
}

func getRouter() *mux.Router {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	deps := initializeDependencies(ctx)
	services := initializeServices(deps)
	handlers := initializeHandlers(services)

	return handlers.InitRouter()
}

type dependencies struct {
	postgres      *gorm.DB
	kafkaConsumer *kafka.Consumer
	ctx           context.Context
}

func initializeDependencies(ctx context.Context) *dependencies {
	postgres, err := config.ConnectPostgres(logger)
	if err != nil {
		logger.Fatal(err)
	}

	kafkaConsumer, err := kafka.NewConsumer("kafka:9092", "user-events", "user-service-group")
	if err != nil {
		logger.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	return &dependencies{
		postgres:      postgres,
		kafkaConsumer: kafkaConsumer,
		ctx:           ctx,
	}
}

func initializeServices(deps *dependencies) *service.Service {
	repo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(deps.postgres),
	)

	userService := user.New(
		user.WithLogger(logger),
		user.WithRepository(repo),
		user.WithKafkaConsumer(deps.kafkaConsumer),
	)

	userService.StartConsumer(deps.ctx)

	return service.New(
		service.WithUserService(userService),
	)
}

func initializeHandlers(services *service.Service) *handler.Handler {
	return handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)
}
