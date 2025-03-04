package routing

import (
	"fmt"
	"github.com/nislovskaya/microservice_architecture/project/route_service/model"
	"github.com/nislovskaya/microservice_architecture/project/route_service/repository"
	"github.com/sirupsen/logrus"
	"time"
)

type Service interface {
	CreateRoute(route *model.Route) error
	GetRoute(id uint) (*model.Route, error)
	UpdateRoute(route *model.Route) error
	DeleteRoute(id uint) error
	GetRoutes() ([]model.Route, error)
	SearchRoutes(from, to, date string) ([]model.Route, error)
}

type routing struct {
	logger *logrus.Entry
	repo   repository.Repository
}

func New(opts ...Option) Service {
	service := &routing{}

	for _, option := range opts {
		option(service)
	}

	return service
}

func (r *routing) CreateRoute(route *model.Route) error {
	if err := r.validateRoute(route); err != nil {
		return err
	}
	return r.repo.CreateRoute(route)
}

func (r *routing) GetRoute(id uint) (*model.Route, error) {
	return r.repo.GetRoute(id)
}

func (r *routing) UpdateRoute(route *model.Route) error {
	if err := r.validateRoute(route); err != nil {
		return err
	}
	return r.repo.UpdateRoute(route)
}

func (r *routing) DeleteRoute(id uint) error {
	return r.repo.DeleteRoute(id)
}

func (r *routing) GetRoutes() ([]model.Route, error) {
	return r.repo.GetRoutes()
}

func (r *routing) SearchRoutes(from, to, dateStr string) ([]model.Route, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	return r.repo.SearchRoutes(from, to, date)
}

func (r *routing) validateRoute(route *model.Route) error {
	if route.StartPoint == "" || route.EndPoint == "" {
		return fmt.Errorf("start point and end point are required")
	}
	if route.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if route.Capacity <= 0 {
		return fmt.Errorf("capacity must be greater than 0")
	}
	return nil
}
