package models

import (
	"time"
)

// struct only

type Users struct {
	Model
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	NamePrefix   string `json:"name_prefix" validate:"required"`
	SumGrade     string `json:"sum_grade" gorm:"default:null"`
	AGrade       string `json:"a_grade" gorm:"default:null"`
	BGrade       string `json:"b_grade" gorm:"default:null"`
	CGrade       string `json:"c_grade" gorm:"default:null"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	ImagePath    string `json:"image_path"`
	ImageID      int    `json:"image_id" gorm:"default:null"`
	CreatedBy    int    `json:"created_by"`
	UpdatedBy    int    `json:"updated_by" gorm:"default:null"`
}

type UserGrade struct {
	ID        int       `json:"id"`
	SumGrade  string    `json:"sum_grade"`
	AGrade    string    `json:"a_grade" validate:"required,oneof=A B C D"`
	BGrade    string    `json:"b_grade" validate:"required,oneof=A B C D"`
	CGrade    string    `json:"c_grade" validate:"required,oneof=A B C D"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=4,lte=6"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserBookings struct {
	ID       int          `json:"id" `
	FullName string       `json:"full_name"`
	Bookings []MyBookings `json:"bookings" gorm:"foreignKey:UsersID;references:ID"`
}

type MyBookings struct {
	UsersID      int    `json:"users_id"`
	BookingTime  string `json:"booking_time"`
	RoomName     string `json:"room_name"`
	AmountPerson int    `json:"amount_person"`
}

type UserWithBookings struct {
	Username     string    `json:"username"`
	Password     string    `json:"password" validate:"required,gte=4,lte=6"`
	StartDate    time.Time `json:"start_date" validate:"required,isBookOverLappingTime"`
	EndDate      time.Time `json:"end_date" validate:"required,isBookOverLappingTime"`
	AmountPerson int       `json:"amount_person" validate:"required,isAmountPersonMoreMaximumPerson" query:"amount_person"`
	RoomsID      int       `json:"rooms_id" validate:"required" query:"rooms_id"`
}

func (MyBookings) TableName() string {
	return "bookings"
}

func (HistoryBookingRooms) TableName() string {
	return "users"
}

func (AmountOfBookingUsers) TableName() string {
	return "users"
}

func (UserGrade) TableName() string {
	return "users"
}
