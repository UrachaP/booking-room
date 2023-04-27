package roomrepository

import (
	"bookingrooms/internal/models"
)

func (r repository) CreateRoomImage(roomImage models.RoomsImages) error {
	return r.db.Create(&roomImage).Error
}
