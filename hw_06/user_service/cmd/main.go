package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nislovskaya/golang_tools/config"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/handler"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/repository"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service/user"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var logger = logrus.NewEntry(logrus.New())

func main() {
	db := config.ConnectDB(logger)

	router := getRouter(db)

	logger.Info("Server is started...")

	logger.Fatal(http.ListenAndServe(":8082", router))
}

func getRouter(db *gorm.DB) *mux.Router {
	userRepository := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(db),
	)

	userService := user.New(
		user.WithLogger(logger),
		user.WithRepository(userRepository),
	)

	services := service.New(
		service.WithUserService(userService),
	)

	handlers := handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)

	return handlers.InitRouter()
}
