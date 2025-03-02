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
	db := config.ConnectPostgres(logger)
	secretKey := config.GetSecret()

	redisClient := config.ConnectRedis(logger)

	router := getRouter(db, secretKey, redisClient)

	logger.Info("Server is started...")
	logger.Fatal(http.ListenAndServe(":8080", router))
}

func getRouter(db *gorm.DB, secretKey string, redisClient *redis.Client) *mux.Router {
	kafkaProducer, err := kafka.NewProducer("kafka:9092")
	if err != nil {
		logger.Fatalf("Failed to create Kafka producer: %v", err)
	}

	repo := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(db),
	)

	authService := auth.New(
		auth.WithLogger(logger),
		auth.WithRepository(repo),
		auth.WithSecretKey(secretKey),
		auth.WithKafkaProducer(kafkaProducer),
		auth.WithRedis(redisClient),
	)

	services := service.New(
		service.WithAuthService(authService),
	)

	handlers := handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)

	return handlers.InitRouter()
}
