package models

import "time"

type Bookings struct {
	ID           int       `json:"id"`
	StartDate    time.Time `json:"start_date" validate:"required,isBookOverLappingTime" query:"start_date"`
	EndDate      time.Time `json:"end_date" validate:"required,isBookOverLappingTime" query:"end_date"`
	AmountPerson int       `json:"amount_person" validate:"required,isAmountPersonMoreMaximumPerson" query:"amount_person"`
	UsersID      int       `json:"users_id" validate:"required" gorm:"users_id" query:"users_id"`
	RoomsID      int       `json:"rooms_id" validate:"required" query:"rooms_id"`
}

type FilterBookings struct {
	FullName     string    `json:"full_name,omitempty" query:"full_name"`
	StartDate    time.Time `json:"start_date,omitempty" validate:"required,isBookOverLappingTime" query:"start_date"`
	EndDate      time.Time `json:"end_date,omitempty" validate:"required,isBookOverLappingTime" query:"end_date"`
	AmountPerson int       `json:"amount_person,omitempty" validate:"required,isAmountPersonMoreMaximumPerson" query:"amount_person"`
	UsersID      int       `json:"users_id,omitempty" validate:"required" gorm:"users_id" query:"users_id"`
	RoomsID      int       `json:"rooms_id,omitempty" validate:"required" query:"rooms_id"`
}

type AmountOfBookingUsers struct {
	FullName      string `json:"full_name"`
	BookingsCount int    `json:"bookings_count"`
}

type AmountOfBookingRooms struct {
	RoomName string `json:"room_name"`
	Count    int    `json:"count"`
}

type HistoryBookingRooms struct {
	FullName      string    `json:"full_name"`
	RoomName      string    `json:"room_name"`
	MaximumPerson int       `json:"maximum_person"`
	AmountPerson  int       `json:"amount_person"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}
