package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/kafka"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/model"
	"github.com/nislovskaya/microservice_architecture/project/booking_service/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateBooking(booking *model.Booking) error
	GetBooking(id uint) (*model.Booking, error)
	UpdateBooking(booking *model.Booking) error
	GetUserBookings(userID uint) ([]model.Booking, error)
	CancelBooking(id uint) error
	ConfirmBooking(id uint) error
}

type booking struct {
	logger   *logrus.Entry
	repo     repository.Repository
	consumer *kafka.Consumer
}

func New(opts ...Option) Service {
	service := &booking{}

	for _, option := range opts {
		option(service)
	}

	return service
}

func (b *booking) CreateBooking(booking *model.Booking) error {
	if err := b.validateBooking(booking); err != nil {
		return err
	}

	booking.Status = "pending"
	return b.repo.CreateBooking(booking)
}

func (b *booking) GetBooking(id uint) (*model.Booking, error) {
	return b.repo.GetBooking(id)
}

func (b *booking) UpdateBooking(booking *model.Booking) error {
	if err := b.validateBooking(booking); err != nil {
		return err
	}
	return b.repo.UpdateBooking(booking)
}

func (b *booking) GetUserBookings(userID uint) ([]model.Booking, error) {
	return b.repo.GetUserBookings(userID)
}

func (b *booking) CancelBooking(id uint) error {
	booking, err := b.repo.GetBooking(id)
	if err != nil {
		return err
	}

	booking.Status = "cancelled"
	return b.repo.UpdateBooking(booking)
}

func (b *booking) ConfirmBooking(id uint) error {
	booking, err := b.repo.GetBooking(id)
	if err != nil {
		return err
	}

	booking.Status = "confirmed"
	return b.repo.UpdateBooking(booking)
}

func (b *booking) validateBooking(booking *model.Booking) error {
	if booking.UserID == 0 {
		return fmt.Errorf("user ID is required")
	}
	if booking.RouteID == 0 {
		return fmt.Errorf("route ID is required")
	}
	if booking.Seats <= 0 {
		return fmt.Errorf("number of seats must be greater than 0")
	}
	return nil
}

func (b *booking) StartConsumer(ctx context.Context) {
	if b.consumer == nil {
		b.logger.Error("Consumer is not initialized")
		return
	}

	go func() {
		b.logger.Info("Starting Kafka consumer for route events")
		if err := b.consumer.Consume(ctx, b.handleMessage); err != nil {
			b.logger.Errorf("Failed to consume message: %v", err)
		}
	}()
}

func (b *booking) handleMessage(message []byte) error {
	var event model.RouteEvent
	if err := json.Unmarshal(message, &event); err != nil {
		return fmt.Errorf("failed to unmarshal route event: %w", err)
	}

	b.logger.Infof("Received route event: %+v", event)

	switch event.Type {
	case "ROUTE_DELETED":
		return b.handleRouteDeleted(event.RouteID)
	case "ROUTE_UPDATED":
		return b.handleRouteUpdated(event)
	}

	return nil
}

func (b *booking) handleRouteDeleted(routeID uint) error {
	bookings, err := b.repo.GetRouteBookings(routeID)
	if err != nil {
		return fmt.Errorf("failed to get route bookings: %w", err)
	}

	for _, booking := range bookings {
		if booking.Status == "pending" {
			booking.Status = "cancelled"
			if err := b.repo.UpdateBooking(&booking); err != nil {
				b.logger.Errorf("Failed to cancel booking %d: %v", booking.ID, err)
			}
		}
	}

	return nil
}

func (b *booking) handleRouteUpdated(event model.RouteEvent) error {
	bookings, err := b.repo.GetRouteBookings(event.RouteID)
	if err != nil {
		return fmt.Errorf("failed to get route bookings: %w", err)
	}

	totalSeats := 0
	for _, booking := range bookings {
		if booking.Status != "cancelled" {
			totalSeats += booking.Seats
		}
	}

	if totalSeats > event.Capacity {
		return b.cancelExcessBookings(bookings, event.Capacity)
	}

	return nil
}

func (b *booking) cancelExcessBookings(bookings []model.Booking, capacity int) error {
	totalSeats := 0
	for _, booking := range bookings {
		if booking.Status == "pending" {
			booking.Status = "cancelled"
			if err := b.repo.UpdateBooking(&booking); err != nil {
				b.logger.Errorf("Failed to cancel booking %d: %v", booking.ID, err)
			}
		} else if booking.Status == "confirmed" {
			totalSeats += booking.Seats
			if totalSeats > capacity {
				booking.Status = "cancelled"
				if err := b.repo.UpdateBooking(&booking); err != nil {
					b.logger.Errorf("Failed to cancel booking %d: %v", booking.ID, err)
				}
			}
		}
	}
	return nil
}
