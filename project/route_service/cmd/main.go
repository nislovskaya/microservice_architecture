package main

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/project/route_service/handler"
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
	router := getRouter()

	logger.Info("Route Service is starting...")
	logger.Fatal(http.ListenAndServe(":8082", router))
}

func getRouter() *mux.Router {
	services := initializeServices()
	handlers := initializeHandlers(services)

	return handlers.InitRouter()
}

func initializeServices() *service.Service {
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
