package bookingrepository

import (
	"strings"
	"time"

	"bookingrooms/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return repository{db: db}
}

type BookingRepository interface {
	CreateBooking(booking models.Bookings) error
	GetBookingById(id int) (models.Bookings, error)
	GetBookings() ([]models.Bookings, error)
	GetCountBookingById(id int) int64
	UpdateBooking(booking models.Bookings) error
	DeleteBooking(booking *models.Bookings) error
	GetBookingFilter(filterBooking models.FilterBookings) (*[]models.Bookings, error)
}

func (r repository) GetCountBookingById(id int) int64 {
	return r.db.First(&models.Bookings{}, id).RowsAffected
}

func (r repository) CreateBooking(booking models.Bookings) error {
	return r.db.Create(&booking).Error
}

func (r repository) GetBookingById(id int) (models.Bookings, error) {
	var booking models.Bookings
	return booking, r.db.First(&booking, id).Error
}

func (r repository) GetBookings() ([]models.Bookings, error) {
	var bookings []models.Bookings
	return bookings, r.db.Find(&bookings).Error
}

func (r repository) UpdateBooking(booking models.Bookings) error {
	return r.db.Updates(&booking).Error
}

func (r repository) DeleteBooking(booking *models.Bookings) error {
	return r.db.Delete(&booking).Error
}

func (r repository) GetBookingFilter(filterBooking models.FilterBookings) (*[]models.Bookings, error) {
	var booking *[]models.Bookings

	preQuery := r.db.Model(&models.Bookings{})

	if filterBooking.FullName != "" {
		fullName := strings.Split(filterBooking.FullName, " ")
		preQuery.Joins("LEFT JOIN users ON users.id = users_id")
		for _, s := range fullName {
			preQuery.Where("first_name like ?", "%"+s+"%").
				Or("last_name like ?", "%"+s+"%")
		}
	}

	nullTime := time.Time{}
	if filterBooking.EndDate != nullTime && filterBooking.StartDate != nullTime {
		preQuery.Where(r.db.
			Where("start_date >= ?", filterBooking.StartDate).
			Where("start_date <= ?", filterBooking.EndDate)).
			Or(r.db.
				Where("end_date >= ?", filterBooking.StartDate).
				Where("end_date <= ?", filterBooking.EndDate))
	}

	if filterBooking.AmountPerson != 0 {
		preQuery.Where("amount_person = ?", filterBooking.AmountPerson)
	}

	if filterBooking.UsersID != 0 {
		preQuery.Where("users_id = ?", filterBooking.UsersID)
	}

	if filterBooking.RoomsID != 0 {
		preQuery.Where("rooms_id = ?", filterBooking.RoomsID)
	}

	return booking, preQuery.Find(booking).Error
}
