package roomservice

import (
	"errors"

	"bookingrooms/internal/models"
	roomrepository "bookingrooms/internal/repositories/room"
)

type service struct {
	repository roomrepository.RoomRepository
}

func NewRoomService(repository roomrepository.RoomRepository) service {
	return service{repository: repository}
}

type RoomService interface {
	SaveRoom(rooms models.Rooms) error
	GetRoom(id int) (models.Rooms, error)
	GetRooms() ([]models.Rooms, error)
	UpdateRoom(room models.Rooms) error
	DeleteRoom(id int) error
	SaveRoomImage(roomImage models.RoomsImages) error
}

func (s service) SaveRoom(rooms models.Rooms) error {
	return s.repository.CreateRoom(rooms)
}

func (s service) GetRoom(id int) (models.Rooms, error) {
	return s.repository.GetRoomById(id)
}

func (s service) GetRooms() ([]models.Rooms, error) {
	return s.repository.GetRooms()
}

func (s service) UpdateRoom(room models.Rooms) error {
	count, err := s.repository.GetCountRoomById(room.ID)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no data")
	}
	return s.repository.UpdateRoom(room)
}

func (s service) DeleteRoom(id int) error {
	return s.repository.DeleteRoom(&models.Rooms{ID: id})
}
