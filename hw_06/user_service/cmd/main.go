package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/handler"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/kafka"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/repository"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/service"
	"github.com/nislovskaya/microservice_architecture/hw_06/user_service/service/user"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
	"net/http"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	kafkaConsumer, err := kafka.NewConsumer("kafka:9092", "user-events", "user-service-group")
	if err != nil {
		logger.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := getRouter(kafkaConsumer, ctx)

	logger.Info("Server is started...")
	logger.Fatal(http.ListenAndServe(":8081", router))
}

func getRouter(consumer *kafka.Consumer, ctx context.Context) *mux.Router {
	services := initializeServices(consumer, ctx)
	handlers := initializeHandlers(services)

	return handlers.InitRouter()
}

func initializeServices(consumer *kafka.Consumer, ctx context.Context) *service.Service {
	postgres, err := config.ConnectPostgres(logger)
	if err != nil {
		logger.Fatal(err)
	}

	repo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(postgres),
	)

	userService := user.New(
		user.WithLogger(logger),
		user.WithRepository(repo),
		user.WithKafkaConsumer(consumer),
	)

	userService.StartConsumer(ctx)

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
