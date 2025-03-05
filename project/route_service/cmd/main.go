package main

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/project/route_service/handler"
	"github.com/nislovskaya/microservice_architecture/project/route_service/kafka"
	"github.com/nislovskaya/microservice_architecture/project/route_service/repository"
	"github.com/nislovskaya/microservice_architecture/project/route_service/service"
	"github.com/nislovskaya/microservice_architecture/project/route_service/service/routing"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
	"net/http"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	kafkaProducer, err := kafka.NewProducer("kafka:9092")
	if err != nil {
		logger.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	router := getRouter(kafkaProducer)

	logger.Info("Route Service is starting...")
	logger.Fatal(http.ListenAndServe(":8082", router))
}

func getRouter(producer *kafka.Producer) *mux.Router {
	services := initializeServices(producer)
	handlers := initializeHandlers(services)

	return handlers.InitRouter()
}

func initializeServices(producer *kafka.Producer) *service.Service {
	postgres, err := config.ConnectPostgres(logger)
	if err != nil {
		logger.Fatal(err)
	}

	routeRepo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(postgres),
	)

	routingService := routing.New(
		routing.WithLogger(logger),
		routing.WithRepository(routeRepo),
		routing.WithKafkaProducer(producer),
	)

	return service.New(
		service.WithRoutingService(routingService),
	)
}

func initializeHandlers(services *service.Service) *handler.Handler {
	return handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)
}
