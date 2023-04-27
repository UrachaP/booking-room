package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"bookingrooms/internal/models"
	"github.com/labstack/echo/v4"
)

func (h Handlers) SaveBooking(c echo.Context) error {
	var booking models.Bookings

	err := c.Bind(&booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.bookingService.CreateBooking(booking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "create booking success")
}

func (h Handlers) GetBooking(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	booking, err := h.bookingService.GetBooking(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, booking)
}

func (h Handlers) GetBookingList(c echo.Context) error {
	bookings, err := h.bookingService.GetBookings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bookings)
}

func (h Handlers) UpdateBooking(c echo.Context) error {
	var booking models.Bookings
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	booking.ID = id

	err = c.Bind(&booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.bookingService.UpdateBooking(booking)
	if err == errors.New("no data") {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "update booking success")
}

func (h Handlers) DeleteBooking(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.bookingService.DeleteBooking(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "deleted booking success")
}

func (h Handlers) GetBookingFilter(c echo.Context) error {
	//query param can use bind
	var filterBooking models.FilterBookings
	err := c.Bind(&filterBooking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	bookings, err := h.bookingService.GetBookingFilter(filterBooking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &bookings)
}
