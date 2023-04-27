package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"bookingrooms/internal/models"
	"github.com/labstack/echo/v4"
)

func (h Handlers) SaveRoom(c echo.Context) error {
	var room models.Rooms

	err := c.Bind(&room)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(room)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.roomService.SaveRoom(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "create room success")
}

func (h Handlers) GetRoom(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	room, err := h.roomService.GetRoom(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, room)
}

func (h Handlers) GetRoomList(c echo.Context) error {
	rooms, err := h.roomService.GetRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rooms)
}

func (h Handlers) UpdateRoom(c echo.Context) error {
	// get id by param
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// read request
	var room models.Rooms
	err = c.Bind(&room)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate
	err = c.Validate(room)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// update room where id
	room.ID = id
	err = h.roomService.UpdateRoom(room)
	if err == errors.New("no data") {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "update room success")
}

func (h Handlers) DeleteRoom(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = h.roomService.DeleteRoom(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "deleted room success")
}

func (h Handlers) SaveRoomImage(c echo.Context) error {
	var roomImage models.RoomsImages
	requestId := c.Param("id")
	err := c.Bind(&roomImage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	roomImage.RoomID = id

	// check have image_id in db
	err = h.tempImageService.HaveTempImages(roomImage.ImageId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = h.roomService.SaveRoomImage(roomImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "update images room success")
}
