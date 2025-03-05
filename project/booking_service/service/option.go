package service

import "github.com/nislovskaya/microservice_architecture/project/booking_service/service/booking"

type Option func(fs *Service)

func WithBookingService(bookingService booking.Service) Option {
	return func(s *Service) {
		s.Service = bookingService
	}
}
