package repository

import (
	"github.com/nislovskaya/microservice_architecture/project/booking_service/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBooking(booking *model.Booking) error
	GetBooking(id uint) (*model.Booking, error)
	UpdateBooking(booking *model.Booking) error
	GetUserBookings(userID uint) ([]model.Booking, error)
	GetRouteBookings(routeID uint) ([]model.Booking, error)
}

type booking struct {
	logger *logrus.Entry
	db     *gorm.DB
}

func New(opts ...Option) Repository {
	repository := &booking{}

	for _, option := range opts {
		option(repository)
	}

	return repository
}

func (b *booking) CreateBooking(booking *model.Booking) error {
	return b.db.Create(booking).Error
}

func (b *booking) GetBooking(id uint) (*model.Booking, error) {
	var booking model.Booking
	if err := b.db.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (b *booking) UpdateBooking(booking *model.Booking) error {
	return b.db.Save(booking).Error
}

func (b *booking) GetUserBookings(userID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	if err := b.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (b *booking) GetRouteBookings(routeID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	if err := b.db.Where("route_id = ?", routeID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
