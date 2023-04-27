package config

import (
	"log"
	"time"

	"bookingrooms/internal/models"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type Handlers struct {
	DB *gorm.DB
}

func (h Handlers) InitValidate() *CustomValidator {
	v := validator.New()
	_ = v.RegisterValidation("isAmountPersonMoreMaximumPerson", h.isAmountPersonMoreMaximumPerson)
	_ = v.RegisterValidation("isBookOverLappingTime", h.isBookOverLappingTime)
	return &CustomValidator{validator: v}
}

func (h Handlers) isAmountPersonMoreMaximumPerson(fl validator.FieldLevel) bool {
	amountPerson := fl.Field().Interface().(int)
	roomId := fl.Parent().FieldByName("RoomsID").Interface().(int)
	var room models.Rooms
	err := h.DB.Where("id = ?", roomId).First(&room).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return amountPerson <= room.MaximumPerson
}

func (h Handlers) isBookOverLappingTime(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(time.Time)
	roomId := fl.Parent().FieldByName("RoomsID").Interface().(int)

	var count int
	err := h.DB.
		Model(&models.Bookings{}).Debug().
		Select("COUNT(*)").
		Where("rooms_id = ?", roomId).
		Where("start_date <= ? and end_date >= ?", date, date).
		First(&count).
		Error
	if err != nil {
		log.Println(err)
		return false
	}
	if count != 0 {
		return false
	}

	return true
}
