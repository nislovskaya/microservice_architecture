package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/handler"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/kafka"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/repository"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/service"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/service/booking"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
	"net/http"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	kafkaConsumer, err := kafka.NewConsumer("kafka:9092", "route-events", "booking-service-group")
	if err != nil {
		logger.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := getRouter(kafkaConsumer, ctx)

	logger.Info("Booking Service is starting...")
	logger.Fatal(http.ListenAndServe(":8083", router))
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

	bookingRepo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(postgres),
	)

	bookingService := booking.New(
		booking.WithLogger(logger),
		booking.WithRepository(bookingRepo),
		booking.WithKafkaConsumer(consumer),
	)

	bookingService.StartConsumer(ctx)

	return service.New(
		service.WithBookingService(bookingService),
	)
}

func initializeHandlers(services *service.Service) *handler.Handler {
	return handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)
}
