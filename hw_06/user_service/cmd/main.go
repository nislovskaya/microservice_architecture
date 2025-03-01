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
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"net/http"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	db := config.ConnectDB(logger)

	kafkaConsumer, err := kafka.NewConsumer("kafka:9092", "user-events", "user-service-group")
	if err != nil {
		logger.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := getRouter(db, kafkaConsumer, ctx)

	logger.Info("Server is started...")

	logger.Fatal(http.ListenAndServe(":8082", router))
}

func getRouter(db *gorm.DB, kafkaConsumer *kafka.Consumer, ctx context.Context) *mux.Router {
	repo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(db),
	)

	userService := user.New(
		user.WithLogger(logger),
		user.WithRepository(repo),
		user.WithKafkaConsumer(kafkaConsumer),
	)

	userService.StartConsumer(ctx)

	services := service.New(
		service.WithUserService(userService),
	)

	handlers := handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)

	return handlers.InitRouter()
}
