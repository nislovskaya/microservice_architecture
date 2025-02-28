package main

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/handler"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/service"
	"github.com/nislovskaya/microservice_architecture/hw_06/jwt_verifier/service/validation"
	"github.com/sirupsen/logrus"
	"net/http"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	secretKey := config.GetSecret()

	router := getRouter(secretKey)

	logger.Info("Server is started...")

	logger.Info("Server is started...")

	logger.Fatal(http.ListenAndServe(":8088", router))
}

func getRouter(secretKey string) *mux.Router {
	validationService := validation.New(
		validation.WithLogger(logger),
		validation.WithSecretKey(secretKey),
	)

	services := service.New(
		service.WithAuthService(validationService),
	)

	handlers := handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)

	return handlers.InitRouter()
}
