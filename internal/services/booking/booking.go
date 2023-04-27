package bookingService

import (
	"errors"

	"bookingrooms/internal/models"
	bookingrepository "bookingrooms/internal/repositories/booking"
)

type service struct {
	repository bookingrepository.BookingRepository
}

func NewBookingService(repository bookingrepository.BookingRepository) service {
	return service{repository: repository}
}

type BookingService interface {
	CreateBookingWithUserId(userWithBooking models.UserWithBookings, userId int) error
	CreateBooking(booking models.Bookings) error
	GetBooking(id int) (models.Bookings, error)
	GetBookings() ([]models.Bookings, error)
	UpdateBooking(booking models.Bookings) error
	DeleteBooking(id int) error
	GetBookingFilter(filterBooking models.FilterBookings) (*[]models.Bookings, error)
}

func (s service) CreateBookingWithUserId(userWithBooking models.UserWithBookings, userId int) error {
	booking := models.Bookings{
		StartDate:    userWithBooking.StartDate,
		EndDate:      userWithBooking.EndDate,
		AmountPerson: userWithBooking.AmountPerson,
		UsersID:      userId,
		RoomsID:      userWithBooking.RoomsID,
	}
	return s.repository.CreateBooking(booking)
}

func (s service) CreateBooking(booking models.Bookings) error {
	return s.repository.CreateBooking(booking)
}

func (s service) GetBooking(id int) (models.Bookings, error) {
	return s.repository.GetBookingById(id)
}

func (s service) GetBookings() ([]models.Bookings, error) {
	return s.repository.GetBookings()
}

func (s service) UpdateBooking(booking models.Bookings) error {
	count := s.repository.GetCountBookingById(booking.ID)
	if count == 0 {
		return errors.New("no data")
	}
	return s.repository.UpdateBooking(booking)
}

func (s service) DeleteBooking(id int) error {
	return s.repository.DeleteBooking(&models.Bookings{ID: id})
}

func (s service) GetBookingFilter(filterBooking models.FilterBookings) (*[]models.Bookings, error) {
	return s.repository.GetBookingFilter(filterBooking)
}
