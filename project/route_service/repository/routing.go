package repository

import (
	"github.com/nislovskaya/microservice_architecture/project/route_service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRoute(route *model.Route) error
	GetRoute(id uint) (*model.Route, error)
	UpdateRoute(route *model.Route) error
	DeleteRoute(id uint) error
	GetRoutes() ([]model.Route, error)
	SearchRoutes(from, to string) ([]model.Route, error)
}

type routing struct {
	logger *logrus.Entry
	db     *gorm.DB
}

func New(opts ...Option) Repository {
	repository := &routing{}

	for _, option := range opts {
		option(repository)
	}

	return repository
}

func (r *routing) CreateRoute(route *model.Route) error {
	return r.db.Create(route).Error
}

func (r *routing) GetRoute(id uint) (*model.Route, error) {
	var route model.Route
	if err := r.db.First(&route, id).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *routing) UpdateRoute(route *model.Route) error {
	return r.db.Save(route).Error
}

func (r *routing) DeleteRoute(id uint) error {
	return r.db.Delete(&model.Route{}, id).Error
}

func (r *routing) GetRoutes() ([]model.Route, error) {
	var routes []model.Route
	if err := r.db.Find(&routes).Error; err != nil {
		return nil, err
	}
	return routes, nil
}

func (r *routing) SearchRoutes(from, to string) ([]model.Route, error) {
	var routes []model.Route
	query := r.db.Where("from_city = ? AND to_city = ? ", from, to)
	if err := query.Find(&routes).Error; err != nil {
		return nil, err
	}
	return routes, nil
}
